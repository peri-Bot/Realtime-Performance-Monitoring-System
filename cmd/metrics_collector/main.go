package main

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/peri-Bot/Realtime-Performance-Monitoring-System/internal/metrics"
)

func main() {
	scheduler := gocron.NewScheduler(time.UTC)

	// Schedule the CollectMetrics function to run every 10 seconds
	scheduler.Every(1).Seconds().Do(runCronJob)

	scheduler.StartBlocking() // Start the scheduler and block the main thread
}

func runCronJob() {
	metric, err := metrics.CollectMetrics()
	if err != nil {
		log.Println("Error collecting metrics:", err)
		return
	}
	metrics.LogMetrics(metric)
}
