package repository

import (
	"candle-collector/internal/config"
	"candle-collector/internal/model"
	"fmt"
)

func GetLastCandleOpenTime(symbol model.Symbol) int64 {
	candle := model.Candle{}
	config.DB.Select("open_time").Where("symbol = ?", symbol.Code).Last(&candle)
	return candle.OpenTime
}

func SaveCandleAll(candles *[]model.Candle) {
	result := config.DB.Create(&candles)
	if result.Error != nil {
		fmt.Println("failed to insert users:", result.Error)
	} else {
		fmt.Printf("%d users inserted successfully.\n", result.RowsAffected)
	}
}
