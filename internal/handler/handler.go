package handler

import (
	"candle-collector/internal/model/symbols"
	"candle-collector/internal/scheduler"
	"net/http"
)

func Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/symbol/collect-start", symbols.SetSymbol)
	mux.HandleFunc("/symbol/collect-stop", symbols.DeleteSymbol)
	mux.HandleFunc("/symbol/collect-past-start", scheduler.PastCandleUpdate)
	return mux
}
