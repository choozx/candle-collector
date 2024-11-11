package test

import (
	"candle-collector/internal/model"
	"candle-collector/internal/repository"
	"testing"
)

func TestGetLastCandleOpenTime(t *testing.T) {
	type args struct {
		symbol model.Symbol
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "BTC",
			args: args{model.BTCUSDT},
			want: 1730967060000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repository.GetLastCandleOpenTime(tt.args.symbol); got != tt.want {
				t.Errorf("GetLastCandleOpenTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
