package test

import (
	"candle-collector/internal/model"
	"candle-collector/internal/service"
	"testing"
)

func TestGetSymbolList(t *testing.T) {
	setup(t)

	type args struct {
		symbol model.Symbol
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{model.ETHUSDT},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			candleService.GetSymbolList(tt.args.symbol)
		})
	}
}
