package metrics

import (
	"fmt"
)

func ProcessMetrics(metrics *Metrics) map[string]interface{} {
	totalCPUUsage := calculateTotalCPUUsage(metrics.CPU)
	return map[string]interface{}{
		"total_cpu_usage": totalCPUUsage,
		"cpu_usage":       formatCPUUsage(metrics.CPU),
		"memory": map[string]interface{}{
			"total": metrics.Memory.Total,
			"used":  metrics.Memory.Used,
			"free":  metrics.Memory.Free,
			"usage": metrics.Memory.Usage,
		},
		"disk": map[string]interface{}{
			"total": metrics.Disk.Total,
			"used":  metrics.Disk.Used,
			"free":  metrics.Disk.Free,
			"usage": metrics.Disk.Usage,
		},
		"timestamp": metrics.Timestamp,
	}
}

// calculateTotalCPUUsage calculates the average CPU usage from the slice of CPU percentages
func calculateTotalCPUUsage(cpuUsage []float64) float64 {
	var total float64
	for _, usage := range cpuUsage {
		total += usage
	}
	fmt.Println("Total CPU: %.2f%%", total)
	return total
}

// formatCPUUsage formats the CPU usage slice into a more readable string
func formatCPUUsage(cpuUsage []float64) string {
	var formatted string
	for i, usage := range cpuUsage {
		formatted += fmt.Sprintf("Core %d: %.2f%% ", i, usage)
	}
	return formatted
}
