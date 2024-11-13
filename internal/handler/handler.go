package handler

import (
	"candle-collector/internal/model/symbols"
	"net/http"
)

func Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/symbol/collect-start", symbols.SetSymbol)
	mux.HandleFunc("/symbol/collect-stop", symbols.DeleteSymbol)
	return mux
}
