package test

import (
	"candle-collector/internal/config"
	"testing"
)

func setup(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		config.CloseDB()
	})
}
