package binance

import (
	"candle-collector/internal/model/symbols"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const BaseUrl = "https://fapi.binance.com/"

const CandlesUrl = "fapi/v1/klines"

func GetCandleList(symbol symbols.Symbol, startTime int64, endTime int64) [][]interface{} {
	u, err := url.Parse(BaseUrl + CandlesUrl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil
	}

	query := u.Query()
	query.Set("symbol", symbol.Name)
	query.Set("interval", "1m")
	query.Set("startTime", strconv.FormatInt(startTime, 10))
	query.Set("endTime", strconv.FormatInt(endTime, 10))

	u.RawQuery = query.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		log.Printf("GET request failed: %v", err)
		return nil // 에러가 발생하면 함수를 종료
	}
	defer func() {
		if resp != nil {
			resp.Body.Close() // resp가 nil이 아니면만 호출
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
	}

	// 2차원 배열로 파싱하기 위해 기본 타입을 사용
	var rawData [][]interface{}
	if err := json.Unmarshal(body, &rawData); err != nil {
		log.Printf("Failed to parse JSON: %v", err)
	}

	return rawData
}
