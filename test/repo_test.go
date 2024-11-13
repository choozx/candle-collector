package test

import (
	"candle-collector/internal/config"
	"candle-collector/internal/model/symbols"
	"testing"
)

func TestSave(t *testing.T) {
	type args struct {
		code     int
		symbol   string
		isUpdate bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				symbol:   "ETHUSDT",
				isUpdate: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			symbol := symbols.Symbol{Name: tt.args.symbol, IsUpdate: tt.args.isUpdate}
			config.DB.Where("name=?", symbol.Name).Save(symbol)
		})
	}
}
