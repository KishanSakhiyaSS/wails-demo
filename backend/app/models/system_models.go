package models

// FinalResponse represents the complete system information response
// @Description Complete system information response containing all system details
type FinalResponse struct {
	CPU      CPU            `json:"cpu"`
	GPUs     []GPU          `json:"gpus"`
	OS       OS             `json:"os"`
	Location Location       `json:"location"`
	Memory   Memory         `json:"memory"`
	Disk     Disk           `json:"disk"`
	Hardware []HardwareInfo `json:"hardware"`
}

// CPU represents CPU information
// @Description CPU information including cores, model, cache, frequency, and usage
type CPU struct {
	Cores     int    `json:"cores" example:"8" description:"Number of CPU cores"`
	Model     string `json:"model" example:"Intel Core i7-10700K" description:"CPU model name"`
	CacheSize string `json:"cache_size" example:"8MB" description:"CPU cache size"`
	Ghz       string `json:"ghz" example:"3.2GHz" description:"CPU frequency"`
	CPUUsage  string `json:"cpu_usage_percentage" example:"45.2%" description:"Current CPU usage percentage"`
}

// GPU represents GPU information
// @Description GPU information including name, VRAM, driver, usage, and clock speed
type GPU struct {
	Name       string `json:"name" example:"NVIDIA GeForce RTX 3080" description:"GPU model name"`
	VRAM       string `json:"vram_size" example:"10GB" description:"GPU video memory size"`
	Driver     string `json:"driver" example:"470.82.01" description:"GPU driver version"`
	Usage      string `json:"usage_percentage" example:"30%" description:"Current GPU usage percentage"`
	ClockSpeed string `json:"clock_speed" example:"1800MHz" description:"GPU clock speed"`
}

// OS represents operating system information
// @Description Operating system information including name, hostname, platform, version, and uptime
type OS struct {
	OS              string `json:"os" example:"linux" description:"Operating system name (e.g., freebsd, linux)"`
	Hostname        string `json:"hostname" example:"my-server" description:"System hostname"`
	Platform        string `json:"platform" example:"ubuntu" description:"Platform name (e.g., ubuntu, linuxmint)"`
	PlatformVersion string `json:"platform_version" example:"20.04.3 LTS" description:"Complete OS version"`
	PlatformFamily  string `json:"platform_family" example:"debian" description:"Platform family (e.g., debian, rhel)"`
	KernelVersion   string `json:"kernel_version" example:"5.4.0-74-generic" description:"OS kernel version"`
	KernelArch      string `json:"kernel_arch" example:"x86_64" description:"Native CPU architecture"`
	Uptime          uint64 `json:"uptime" example:"86400" description:"System uptime in seconds"`
}

// Location represents location information
// @Description Location information including IP, hostname, city, region, country, and timezone
type Location struct {
	Ip       string `json:"ip" example:"192.168.1.100" description:"IP address"`
	Hostname string `json:"hostname" example:"my-server" description:"System hostname"`
	City     string `json:"city" example:"San Francisco" description:"City name"`
	Region   string `json:"region" example:"California" description:"Region/state name"`
	Country  string `json:"country" example:"US" description:"Country code"`
	LOC      string `json:"loc" example:"37.7749,-122.4194" description:"Latitude and longitude coordinates"`
	ORG      string `json:"org" example:"AS1234 Example Corp" description:"Organization/ISP"`
	Postal   string `json:"postal" example:"94105" description:"Postal code"`
	Timezone string `json:"timezone" example:"America/Los_Angeles" description:"Timezone"`
	Readme   string `json:"readme" example:"https://ipinfo.io/missingauth" description:"Additional information URL"`
}

// Memory represents memory information
// @Description Memory information including total, used, free, available, and usage percentage
type Memory struct {
	Total       string `json:"total" example:"16GB" description:"Total memory"`
	Used        string `json:"used" example:"8GB" description:"Used memory"`
	Free        string `json:"free" example:"4GB" description:"Free memory"`
	Available   string `json:"available" example:"12GB" description:"Available memory"`
	UsedPercent string `json:"used_percentage" example:"50%" description:"Memory usage percentage"`
}

// Disk represents disk information
// @Description Disk information including total, used, free, and usage percentage
type Disk struct {
	Total       string `json:"total" example:"1TB" description:"Total disk space"`
	Used        string `json:"used" example:"500GB" description:"Used disk space"`
	Free        string `json:"free" example:"500GB" description:"Free disk space"`
	UsedPercent string `json:"used_percentage" example:"50%" description:"Disk usage percentage"`
}

// HardwareInfo represents hardware/network interface information
// @Description Hardware/network interface information including index, MTU, name, hardware address, and flags
type HardwareInfo struct {
	Index        int    `json:"index" example:"1" description:"Interface index (positive integer)"`
	MTU          int    `json:"mtu" example:"1500" description:"Maximum transmission unit"`
	Name         string `json:"name" example:"eth0" description:"Interface name (e.g., en0, lo0, eth0.100)"`
	HardwareAddr string `json:"hardware_addr" example:"00:11:22:33:44:55" description:"IEEE MAC-48, EUI-48 and EUI-64 form"`
	Flags        string `json:"flags" example:"up,broadcast,running,multicast" description:"Interface flags (e.g., FlagUp, FlagLoopback, FlagMulticast)"`
}

// APIResponse represents a standard API response
// @Description Standard API response structure
type APIResponse struct {
	Status  string `json:"status" example:"ok" description:"Response status"`
	Message string `json:"message,omitempty" example:"System benchmark API is running" description:"Response message"`
	Data    any    `json:"data,omitempty" description:"Response data"`
	Error   string `json:"error,omitempty" description:"Error message if any"`
}

// ErrorResponse represents an error response
// @Description Error response structure
type ErrorResponse struct {
	Error   string `json:"error" example:"Failed to get system information" description:"Error message"`
	Details string `json:"details" example:"connection timeout" description:"Error details"`
}

// SystemInfo represents the complete system information (alias for FinalResponse)
// @Description Complete system information response
type SystemInfo = FinalResponse

// CPUInfo represents CPU information (alias for CPU)
// @Description CPU information
type CPUInfo = CPU

// GPUInfo represents GPU information (alias for GPU)
// @Description GPU information
type GPUInfo = GPU

// OSInfo represents OS information (alias for OS)
// @Description Operating system information
type OSInfo = OS

// LocationInfo represents location information (alias for Location)
// @Description Location information
type LocationInfo = Location

// MemoryInfo represents memory information (alias for Memory)
// @Description Memory information
type MemoryInfo = Memory

// DiskInfo represents disk information (alias for Disk)
// @Description Disk information
type DiskInfo = Disk

// UsagePercentages represents usage percentages for various system components
// @Description Usage percentages for CPU, GPU, memory, and disk
type UsagePercentages struct {
	CPU    string `json:"cpu_usage" example:"45.2%" description:"CPU usage percentage"`
	GPU    string `json:"gpu_usage" example:"30%" description:"GPU usage percentage"`
	Memory string `json:"memory_usage" example:"50%" description:"Memory usage percentage"`
	Disk   string `json:"disk_usage" example:"75%" description:"Disk usage percentage"`
}
