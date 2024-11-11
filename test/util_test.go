package test

import (
	"candle-collector/internal/utils"
	"testing"
	"time"
)

func TestBetweenMinuit(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				utils.NowMilliSecond(),
				utils.Minus(utils.NowMilliSecond(), time.Minute),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.BetweenMinuit(tt.args.a+59999, tt.args.b)
		})
	}
}
