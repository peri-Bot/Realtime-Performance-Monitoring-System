package metrics

import (
	"encoding/json"
	"log"
)

// LogMetrics logs metrics as a JSON string.
func LogMetrics(metrics *Metrics) {

	processedMetrics := ProcessMetrics(metrics)
	data, err := json.MarshalIndent(processedMetrics, "", "  ")
	if err != nil {
		log.Println("Error serializing metrics:", err)
		return
	}
	log.Println(string(data))
}
