package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/kishansakhiya/wails-demo/backend/app/models"
	"github.com/kishansakhiya/wails-demo/backend/app/utils"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

// SystemService handles all system information gathering
type SystemService struct {
}

// NewSystemService creates a new instance of SystemService
func NewSystemService() *SystemService {
	return &SystemService{}
}

// GetAllSystemInfo retrieves all system information
func (s *SystemService) GetAllSystemInfo() (*models.FinalResponse, error) {
	// Use context with timeout for the entire operation
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create channels for concurrent data fetching
	type result struct {
		data any
		err  error
		key  string
	}

	results := make(chan result, 7) // 7 different data types

	// Fetch all data concurrently
	go func() {
		data, err := s.fetchCPUInfo()
		results <- result{data: data, err: err, key: "cpu"}
	}()

	go func() {
		data, err := utils.GetGPUInfo()
		results <- result{data: data, err: err, key: "gpu"}
	}()

	go func() {
		data, err := s.fetchOSInfo()
		results <- result{data: data, err: err, key: "os"}
	}()

	go func() {
		location, err := utils.GetLocationInfo()
		if err != nil {
			results <- result{data: nil, err: err, key: "location"}
		} else {
			results <- result{data: &location, err: nil, key: "location"}
		}
	}()

	go func() {
		data, err := s.fetchMemoryInfo()
		results <- result{data: data, err: err, key: "memory"}
	}()

	go func() {
		data, err := s.fetchDiskInfo()
		results <- result{data: data, err: err, key: "disk"}
	}()

	go func() {
		data, err := s.fetchHardwareInfo()
		results <- result{data: data, err: err, key: "hardware"}
	}()

	// Collect results
	var cpuInfo *models.CPU
	var gpuInfo []models.GPU
	var osInfo *models.OS
	var locationInfo *models.Location
	var memoryInfo *models.Memory
	var diskInfo *models.Disk
	var hardwareInfo []models.HardwareInfo

	for i := 0; i < 7; i++ {
		select {
		case res := <-results:
			if res.err != nil {
				return nil, fmt.Errorf("failed to get %s info: %w", res.key, res.err)
			}
			switch res.key {
			case "cpu":
				cpuInfo = res.data.(*models.CPU)
			case "gpu":
				gpuInfo = res.data.([]models.GPU)
			case "os":
				osInfo = res.data.(*models.OS)
			case "location":
				locationInfo = res.data.(*models.Location)
			case "memory":
				memoryInfo = res.data.(*models.Memory)
			case "disk":
				diskInfo = res.data.(*models.Disk)
			case "hardware":
				hardwareInfo = res.data.([]models.HardwareInfo)
			}
		case <-ctx.Done():
			return nil, fmt.Errorf("timeout while gathering system information")
		}
	}

	return &models.FinalResponse{
		CPU:      *cpuInfo,
		GPUs:     gpuInfo,
		OS:       *osInfo,
		Location: *locationInfo,
		Memory:   *memoryInfo,
		Disk:     *diskInfo,
		Hardware: hardwareInfo,
	}, nil
}

// GetCPUInfo retrieves CPU information
func (s *SystemService) GetCPUInfo() (*models.CPU, error) {
	return s.fetchCPUInfo()
}

// fetchCPUInfo performs the actual CPU information fetching
func (s *SystemService) fetchCPUInfo() (*models.CPU, error) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, fmt.Errorf("failed to get CPU info: %w", err)
	}
	cpuPercent, _ := cpu.Percent(time.Second, false)
	if len(cpuInfo) == 0 {
		return nil, fmt.Errorf("no CPU information available")
	}

	cpu := cpuInfo[0]

	// Get CPU usage - simplified approach for now
	usage := cpuPercent[0]

	return &models.CPU{
		Cores:     cpuInfo[0].Cores,
		Model:     cpu.ModelName,
		CacheSize: fmt.Sprintf("%dMB", cpu.CacheSize/1024),
		Ghz:       fmt.Sprintf("%.2fGHz", float64(cpu.Mhz)/1000),
		CPUUsage:  fmt.Sprintf("%.2f%%", usage),
	}, nil
}

// GetGPUInfo retrieves GPU information
func (s *SystemService) GetGPUInfo() ([]models.GPU, error) {
	return utils.GetGPUInfo()
}

// GetOSInfo retrieves operating system information
func (s *SystemService) GetOSInfo() (*models.OS, error) {
	return s.fetchOSInfo()
}

// fetchOSInfo performs the actual OS information fetching
func (s *SystemService) fetchOSInfo() (*models.OS, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return nil, fmt.Errorf("failed to get host info: %w", err)
	}

	return &models.OS{
		OS:              hostInfo.OS,
		Hostname:        hostInfo.Hostname,
		Platform:        hostInfo.Platform,
		PlatformVersion: hostInfo.PlatformVersion,
		PlatformFamily:  hostInfo.PlatformFamily,
		KernelVersion:   hostInfo.KernelVersion,
		KernelArch:      hostInfo.KernelArch,
		Uptime:          hostInfo.Uptime,
	}, nil
}

// GetLocationInfo retrieves location information
func (s *SystemService) GetLocationInfo() (*models.Location, error) {
	location, err := utils.GetLocationInfo()
	if err != nil {
		return nil, err
	}
	return &location, nil
}

// GetMemoryInfo retrieves memory information
func (s *SystemService) GetMemoryInfo() (*models.Memory, error) {
	return s.fetchMemoryInfo()
}

// fetchMemoryInfo performs the actual memory information fetching
func (s *SystemService) fetchMemoryInfo() (*models.Memory, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("failed to get memory info: %w", err)
	}

	return &models.Memory{
		Total:       utils.FormatBytes(memory.Total, 1000),
		Used:        utils.FormatBytes(memory.Used, 1000),
		Free:        utils.FormatBytes(memory.Free, 1000),
		Available:   utils.FormatBytes(memory.Available, 1000),
		UsedPercent: fmt.Sprintf("%.1f%%", memory.UsedPercent),
	}, nil
}

// GetDiskInfo retrieves disk information
func (s *SystemService) GetDiskInfo() (*models.Disk, error) {
	return s.fetchDiskInfo()
}

// fetchDiskInfo performs the actual disk information fetching
func (s *SystemService) fetchDiskInfo() (*models.Disk, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, fmt.Errorf("failed to get disk partitions: %w", err)
	}

	var totalSize, totalUsed, totalFree uint64

	for _, partition := range partitions {
		// Skip certain filesystem types that might cause issues
		if strings.Contains(strings.ToLower(partition.Fstype), "tmpfs") ||
			strings.Contains(strings.ToLower(partition.Fstype), "devtmpfs") {
			continue
		}

		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			continue
		}

		totalSize += usage.Total
		totalUsed += usage.Used
		totalFree += usage.Free
	}

	usedPercent := float64(0)
	if totalSize > 0 {
		usedPercent = float64(totalUsed) / float64(totalSize) * 100
	}

	return &models.Disk{
		Total:       utils.FormatBytes(totalSize, 1000),
		Used:        utils.FormatBytes(totalUsed, 1000),
		Free:        utils.FormatBytes(totalFree, 1000),
		UsedPercent: fmt.Sprintf("%.1f%%", usedPercent),
	}, nil
}

// GetHardwareInfo retrieves hardware/network interface information
func (s *SystemService) GetHardwareInfo() ([]models.HardwareInfo, error) {
	return s.fetchHardwareInfo()
}

// fetchHardwareInfo performs the actual hardware information fetching
func (s *SystemService) fetchHardwareInfo() ([]models.HardwareInfo, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get network interfaces: %w", err)
	}

	var hardwareInfo []models.HardwareInfo
	for _, iface := range interfaces {
		// Skip loopback interfaces and interfaces without hardware address
		if strings.Contains(strings.ToLower(iface.Name), "lo") ||
			strings.HasPrefix(strings.ToLower(iface.Name), "br-") ||
			strings.HasPrefix(strings.ToLower(iface.Name), "docker") ||
			strings.HasPrefix(strings.ToLower(iface.Name), "veth") ||
			strings.HasPrefix(strings.ToLower(iface.Name), "tun") ||
			strings.HasPrefix(strings.ToLower(iface.Name), "tap") ||
			strings.HasPrefix(strings.ToLower(iface.Name), "wg") ||
			strings.HasPrefix(strings.ToLower(iface.Name), "vnet") ||
			strings.HasPrefix(strings.ToLower(iface.Name), "veth") ||
			iface.HardwareAddr == "" {
			continue
		}

		flags := []string{}
		// Simplified flag checking since gopsutil v3 has different flag constants
		if strings.Contains(strings.ToLower(iface.Name), "up") {
			flags = append(flags, "up")
		}
		if strings.Contains(strings.ToLower(iface.Name), "broadcast") {
			flags = append(flags, "broadcast")
		}
		if strings.Contains(strings.ToLower(iface.Name), "multicast") {
			flags = append(flags, "multicast")
		}

		hardwareInfo = append(hardwareInfo, models.HardwareInfo{
			Index:        iface.Index,
			MTU:          iface.MTU,
			Name:         iface.Name,
			HardwareAddr: iface.HardwareAddr,
			Flags:        strings.Join(flags, "|"),
		})
	}

	return hardwareInfo, nil
}

// GetUsagePercentages retrieves usage percentages for CPU, GPU, memory, and disk
func (s *SystemService) GetUsagePercentages() (*models.UsagePercentages, error) {
	return s.fetchUsagePercentages()
}

// fetchUsagePercentages performs the actual usage percentages fetching
func (s *SystemService) fetchUsagePercentages() (*models.UsagePercentages, error) {
	// Get CPU usage
	cpuInfo, err := s.GetCPUInfo()
	if err != nil {
		return nil, fmt.Errorf("failed to get CPU usage: %w", err)
	}

	// Get GPU usage
	gpuInfo, err := s.GetGPUInfo()
	if err != nil {
		return nil, fmt.Errorf("failed to get GPU usage: %w", err)
	}

	// Get memory usage
	memoryInfo, err := s.GetMemoryInfo()
	if err != nil {
		return nil, fmt.Errorf("failed to get memory usage: %w", err)
	}

	// Get disk usage
	diskInfo, err := s.GetDiskInfo()
	if err != nil {
		return nil, fmt.Errorf("failed to get disk usage: %w", err)
	}

	// Get GPU usage percentage (use the first GPU if available)
	gpuUsage := "0%"
	if len(gpuInfo) > 0 {
		gpuUsage = gpuInfo[0].Usage
	}

	return &models.UsagePercentages{
		CPU:    cpuInfo.CPUUsage,
		GPU:    gpuUsage,
		Memory: memoryInfo.UsedPercent,
		Disk:   diskInfo.UsedPercent,
	}, nil
}
