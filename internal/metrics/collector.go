package metrics

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"time"
)

type ResourceUsage struct {
	Total uint64  `json:"total"`
	Used  uint64  `json:"used"`
	Free  uint64  `json:"free"`
	Usage float64 `json:"usage"` // Percentage usage
}

type Metrics struct {
	CPU       []float64     `json:"cpu_usage"` // CPU usage as percentage
	Memory    ResourceUsage `json:"memory"`
	Disk      ResourceUsage `json:"disk"`
	Timestamp time.Time     `json:"timestamp"`
}

func collectCPU() ([]float64, error) {
	cpuPercentages, err := cpu.Percent(0, true)
	if err != nil {
		return nil, err
	}
	return cpuPercentages, nil // Return all core percentages
}

func collectMemory() (ResourceUsage, error) {
	memStats, err := mem.VirtualMemory()
	if err != nil {
		return ResourceUsage{}, err
	}
	return ResourceUsage{
		Total: memStats.Total,
		Used:  memStats.Used,
		Free:  memStats.Available,
		Usage: memStats.UsedPercent,
	}, nil
}

func collectDisk() (ResourceUsage, error) {
	diskStats, err := disk.Usage("/")
	if err != nil {
		return ResourceUsage{}, err
	}
	return ResourceUsage{
		Total: diskStats.Total,
		Used:  diskStats.Used,
		Free:  diskStats.Free,
		Usage: diskStats.UsedPercent,
	}, nil
}

// CollectMetrics collects all metrics and returns a Metrics struct.
func CollectMetrics() (*Metrics, error) {
	cpu, err := collectCPU()
	if err != nil {
		return nil, err
	}
	mem, err := collectMemory()
	if err != nil {
		return nil, err
	}
	disk, err := collectDisk()
	if err != nil {
		return nil, err
	}
	return &Metrics{
		CPU:       cpu,
		Memory:    mem,
		Disk:      disk,
		Timestamp: time.Now(),
	}, nil
}

func PrintMetrics(m *Metrics) {
	if m == nil {
		fmt.Println("No metrics available")
		return
	}

	fmt.Printf("\nSystem Metrics as of %s\n", m.Timestamp.Format(time.RFC1123))
	fmt.Println("----------------------------------------")

	// CPU
	fmt.Printf("CPU Usage: %.2f%%\n\n", m.CPU)

	// Memory
	fmt.Println("Memory:")
	fmt.Printf("  Total: %s\n", formatBytes(m.Memory.Total))
	fmt.Printf("  Used:  %s\n", formatBytes(m.Memory.Used))
	fmt.Printf("  Free:  %s\n", formatBytes(m.Memory.Free))
	fmt.Printf("  Usage: %.2f%%\n\n", m.Memory.Usage)

	// Disk
	fmt.Println("Disk:")
	fmt.Printf("  Total: %s\n", formatBytes(m.Disk.Total))
	fmt.Printf("  Used:  %s\n", formatBytes(m.Disk.Used))
	fmt.Printf("  Free:  %s\n", formatBytes(m.Disk.Free))
	fmt.Printf("  Usage: %.2f%%\n", m.Disk.Usage)
}

// formatBytes converts bytes to human readable string in KB, MB, GB, or TB
func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
