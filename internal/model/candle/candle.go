package candle

import (
	"candle-collector/internal/model/symbols"
	"strconv"
)

type Candle struct {
	OpenTime                 int64   `gorm:"column:open_time"`
	Symbol                   int     `gorm:"column:symbol"`
	Open                     float64 `gorm:"column:open_price"`
	High                     float64 `gorm:"column:high_price"`
	Low                      float64 `gorm:"column:low_price"`
	Close                    float64 `gorm:"column:close_price"`
	Volume                   float64 `gorm:"column:volume"`
	CloseTime                int64   `gorm:"column:close_time"`
	QuoteAssetVolume         float64 `gorm:"column:quote_asset_volume"`
	NumberOfTrades           int     `gorm:"column:number_of_trades"`
	TakerBuyBaseAssetVolume  float64 `gorm:"column:taker_buy_base_asset_volume"`
	TakerBuyQuoteAssetVolume float64 `gorm:"column:taker_buy_quote_asset_volume"`
}

func NewCandle(data []interface{}, symbol symbols.Symbol) *Candle {
	open, _ := strconv.ParseFloat(data[1].(string), 64)
	high, _ := strconv.ParseFloat(data[2].(string), 64)
	low, _ := strconv.ParseFloat(data[3].(string), 64)
	closePrice, _ := strconv.ParseFloat(data[4].(string), 64)
	volume, _ := strconv.ParseFloat(data[5].(string), 64)
	quoteAssetVolume, _ := strconv.ParseFloat(data[7].(string), 64)
	takerBuyBaseAssetVolume, _ := strconv.ParseFloat(data[9].(string), 64)
	takerBuyQuoteAssetVolume, _ := strconv.ParseFloat(data[10].(string), 64)

	candle := Candle{
		OpenTime:                 int64(data[0].(float64)),
		Symbol:                   symbol.Code,
		Open:                     open,
		High:                     high,
		Low:                      low,
		Close:                    closePrice,
		Volume:                   volume,
		CloseTime:                int64(data[6].(float64)),
		QuoteAssetVolume:         quoteAssetVolume,
		NumberOfTrades:           int(data[8].(float64)),
		TakerBuyBaseAssetVolume:  takerBuyBaseAssetVolume,
		TakerBuyQuoteAssetVolume: takerBuyQuoteAssetVolume,
	}
	return &candle
}

func (Candle) TableName() string {
	return "candle"
}
