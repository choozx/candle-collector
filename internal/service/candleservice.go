package candleService

import (
	"candle-collector/internal/model"
	"candle-collector/internal/repository"
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

func CandleUpdate() {
	for _, symbol := range model.Symbols {
		GetSymbolList(symbol)
	}
}

func GetSymbolList(symbol model.Symbol) {
	now := utils.NowMilliSecond()
	lastCandleOpenTime := repository.GetLastCandleOpenTime(symbol)
	if lastCandleOpenTime == 0 {
		lastCandleOpenTime = utils.Minus(now, (model.MaxUpdateCandleCount+2)*time.Minute)
	}
	startTime := utils.Add(lastCandleOpenTime, time.Minute)
	endTime := utils.Minus(now, time.Minute)

	updateCandleCount := utils.BetweenMinuit(startTime, endTime)
	if updateCandleCount > model.MaxUpdateCandleCount {
		// 최대갯수를 딱 맞추기 위해서는 response에 시작시간 캔들도 포함이기 때문에 1을 빼줘야됨
		endTime = utils.Add(startTime, (model.MaxUpdateCandleCount-1)*time.Minute)
	}

	u, err := url.Parse(model.BaseUrl + model.CandlesUrl)
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
	var candles []model.Candle
	for _, data := range rawData {
		candle := model.NewCandle(data, symbol)
		candles = append(candles, *candle)
	}

	repository.SaveCandleAll(&candles)
}
