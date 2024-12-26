package main

import (
	"github.com/peri-Bot/Realtime-Performance-Monitoring-System/internal/metrics"
	"time"
)

func main() {
	for {
		metrics.CollectMem()
		time.Sleep(5 * time.Second) // Collect metrics every 5 seconds
	}
}
