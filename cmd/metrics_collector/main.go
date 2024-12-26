package main

import (
	"log"
	"time"

	"github.com/peri-Bot/Realtime-Performance-Monitoring-System/internal/metrics"
)

func main() {
	for {
		metric, err := metrics.CollectMetrics()
		if err != nil {
			log.Println("Error collecting metrics:", err)
			continue
		}

		metrics.PrintMetrics(metric)

		time.Sleep(5 * time.Second) // Collect metrics every 5 seconds
	}
}
