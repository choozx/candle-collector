package main

import (
	"candle-collector/internal/config"
	"candle-collector/internal/handler"
	"candle-collector/internal/scheduler"
	"fmt"
	"github.com/robfig/cron/v3"
	"net/http"
)

var c *cron.Cron

func init() {
	c := cron.New(cron.WithSeconds())

	entryID, err := c.AddFunc("2 * * * * *", scheduler.CandleUpdate)
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

	http.ListenAndServe(":8081", handler.Handler())
}
