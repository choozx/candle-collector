package test

import (
	"candle-collector/internal/config"
	"testing"
)

func setup(t *testing.T) {
	t.Helper()
	config.InitDB()
	t.Cleanup(func() {
		config.CloseDB()
	})
}
