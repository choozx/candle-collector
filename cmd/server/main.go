package main

import (
	"candle-collector/internal/config"
	"candle-collector/internal/handler"
	candleService "candle-collector/internal/service"
	"fmt"
	"github.com/robfig/cron/v3"
	"net/http"
)

var c *cron.Cron

func init() {
	config.InitDB()

	c := cron.New(cron.WithSeconds())

	entryID, err := c.AddFunc("2 * * * * *", candleService.CandleUpdate)
	if err != nil {
		fmt.Println("Failed to add cron job:", err)
		return
	}

	fmt.Println("Cron job added with Entry ID:", entryID)
	c.Start()
}

func main() {
	defer c.Stop()
	defer config.CloseDB()

	http.ListenAndServe(":8080", handler.Handler())
}
