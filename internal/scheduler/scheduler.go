package scheduler

import (
	"candle-collector/internal/config"
	"candle-collector/internal/model/binance"
	"candle-collector/internal/model/candle"
	"candle-collector/internal/model/symbols"
	"candle-collector/internal/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const MaxUpdateCandleCount = 500

func CandleUpdate() {
	for _, symbol := range symbols.Symbols {
		GetSymbolList(symbol)
	}
	fmt.Println("캔들 수집 완료")
}

func GetSymbolList(symbol symbols.Symbol) {
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

	u, err := url.Parse(binance.BaseUrl + binance.CandlesUrl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	query := u.Query()
	query.Set("symbol", symbol.Name)
	query.Set("interval", "1m")
	query.Set("startTime", strconv.FormatInt(startTime, 10))
	query.Set("endTime", strconv.FormatInt(endTime, 10))

	u.RawQuery = query.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatalf("GET request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// 2차원 배열로 파싱하기 위해 기본 타입을 사용
	var rawData [][]interface{}
	if err := json.Unmarshal(body, &rawData); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// 배열을 구조체로 변환
	var candles []candle.Candle
	for _, data := range rawData {
		newCandle := candle.NewCandle(data, symbol)
		candles = append(candles, *newCandle)
	}

	saveCandleAll(&candles)
}

// repo ##########
// repo go의 package레이아웃을 잘 모르니끼 분리해야하면 그때 분리하자

func getLastCandleOpenTime(symbol symbols.Symbol) int64 {
	newCandle := candle.Candle{}
	config.DB.Select("open_time").Where("symbol = ?", symbol.Code).Last(&newCandle)
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
