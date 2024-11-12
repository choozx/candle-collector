package test

import (
	"candle-collector/internal/model/candle"
	"candle-collector/internal/repository/candlerepository"
	"testing"
)

func TestGetSymbolList(t *testing.T) {
	setup(t)

	type args struct {
		symbol candle.Symbol
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{candle.ETHUSDT},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			candle.GetSymbolList(tt.args.symbol)
		})
	}
}

func TestGetLastCandleOpenTime(t *testing.T) {
	type args struct {
		symbol candle.Symbol
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "BTC",
			args: args{candle.BTCUSDT},
			want: 1730967060000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candlerepository.GetLastCandleOpenTime(tt.args.symbol); got != tt.want {
				t.Errorf("getLastCandleOpenTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
