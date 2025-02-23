package scheduler

import (
	"candle-collector/internal/config"
	"candle-collector/internal/model/binance"
	"candle-collector/internal/model/candle"
	"candle-collector/internal/model/symbols"
	"candle-collector/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const MaxUpdateCandleCount = 500
const MinOpenDateTime = "2022-01-01 00:00:00"
const YyyyMmDdHhMmSs = "2006-01-02 15:04:05"

func CandleUpdate() {
	for _, symbol := range symbols.Symbols {
		UpdateSymbolList(symbol)
	}
	fmt.Println("캔들 수집 완료")
}

func PastCandleUpdate(writer http.ResponseWriter, request *http.Request) {
	requestSymbol := new(symbols.Symbol)
	err := json.NewDecoder(request.Body).Decode(requestSymbol)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(writer, err)
		return
	}

	if requestSymbol.Name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	symbol := symbols.FindSymbol(requestSymbol.Name)

	go UpdatePastCandle(*symbol)
}

func UpdateSymbolList(symbol symbols.Symbol) {
	now := utils.NowMilliSecond()
	lastCandleOpenTime := getLastCandleOpenTime(symbol)
	if lastCandleOpenTime == 0 {
		lastCandleOpenTime = utils.Minus(now, (MaxUpdateCandleCount+2)*time.Minute)
	}
	startTime := utils.Add(lastCandleOpenTime, time.Minute)
	endTime := utils.Minus(now, time.Minute)

	updateCandleCount := utils.BetweenMinuit(startTime, endTime)
	if updateCandleCount > MaxUpdateCandleCount {
		// 최대갯수를 딱 맞추기 위해서는 response에 시작시간 캔들도 포함이기 때문에 1을 빼줘야됨
		endTime = utils.Add(startTime, (MaxUpdateCandleCount-1)*time.Minute)
	}

	var rawData = binance.GetCandleList(symbol, startTime, endTime)

	// 배열을 구조체로 변환
	var candles []candle.Candle
	for _, data := range rawData {
		newCandle := candle.NewCandle(data, symbol)
		candles = append(candles, *newCandle)
	}

	saveCandleAll(&candles)
}

func UpdatePastCandle(symbol symbols.Symbol) {
	t, err := time.Parse(YyyyMmDdHhMmSs, MinOpenDateTime)
	if err != nil {
		fmt.Println("시간 변환 오류:", err)
		return
	}

	for {
		firstCandleOpenTime := getFirstCandleOpenTime(symbol) // 가장 오래된 캔들 오픈시간
		if firstCandleOpenTime == 0 {
			log.Printf("수집된적 없는 심볼입니다!")
			return
		}

		if firstCandleOpenTime < t.UnixMilli() {
			log.Printf("과거 캔들은 전부 수집된 심볼입니다.")
			return
		}

		endTime := utils.Minus(firstCandleOpenTime, time.Minute)
		startTime := utils.Minus(firstCandleOpenTime, time.Minute*MaxUpdateCandleCount)

		var rawData = binance.GetCandleList(symbol, startTime, endTime)

		// 배열을 구조체로 변환
		var candles []candle.Candle
		for _, data := range rawData {
			newCandle := candle.NewCandle(data, symbol)
			candles = append(candles, *newCandle)
		}

		saveCandleAll(&candles)

		newCandleSize := len(candles)
		log.Printf("%v ~ %v 사이즈:%v", time.UnixMilli(candles[0].OpenTime), time.UnixMilli(candles[newCandleSize-1].OpenTime), newCandleSize)
		time.Sleep(time.Minute) // 5초간 대기
	}
}

// repo ##########
// repo go의 package레이아웃을 잘 모르니끼 분리해야하면 그때 분리하자

func getLastCandleOpenTime(symbol symbols.Symbol) int64 {
	newCandle := candle.Candle{}
	config.DB.Select("open_time").Where("symbol = ?", symbol.Code).Last(&newCandle)
	return newCandle.OpenTime
}

func getFirstCandleOpenTime(symbol symbols.Symbol) int64 {
	newCandle := candle.Candle{}
	config.DB.Select("open_time").Where("symbol = ?", symbol.Code).Order("open_time ASC").First(&newCandle)
	return newCandle.OpenTime
}

func saveCandleAll(candles *[]candle.Candle) {
	result := config.DB.Create(&candles)
	if result.Error != nil {
		fmt.Println("failed to insert users:", result.Error)
	} else {
		fmt.Printf("%d users inserted successfully.\n", result.RowsAffected)
	}
}
