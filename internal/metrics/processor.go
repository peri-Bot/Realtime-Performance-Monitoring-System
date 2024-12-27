package metrics

func ProcessMetrics(metrics *Metrics) map[string]interface{} {
	return map[string]interface{}{
		"cpu_usage": metrics.CPU,
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
