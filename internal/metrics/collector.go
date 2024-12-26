package metrics

import (
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

type ResourceUsage struct {
	Total uint64  `json:"total"`
	Used  uint64  `json:"used"`
	Free  uint64  `json:"free"`
	Usage float64 `json:"usage"` // Percentage usage
}

type Metrics struct {
	CPU       float64       `json:"cpu_usage"` // CPU usage as percentage
	Memory    ResourceUsage `json:"memory"`
	Disk      ResourceUsage `json:"disk"`
	Timestamp time.Time     `json:"timestamp"`
}

func collectCPU() (float64, error) {
	cpuPercentages, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	return cpuPercentages[0], nil
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
