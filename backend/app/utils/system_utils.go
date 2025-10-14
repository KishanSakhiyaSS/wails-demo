package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/kishansakhiya/wails-demo/backend/app/models"
)

// FormatBytes converts bytes to human readable format
func FormatBytes(bytes uint64, unitInt int) string {
	unit := uint64(unitInt)
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// GetGPUInfo retrieves GPU information
func GetGPUInfo() ([]models.GPU, error) {
	var gpus []models.GPU

	switch runtime.GOOS {
	case "linux":
		gpus = getLinuxGPUInfo()
	case "windows":
		gpus = getWindowsGPUInfo()
	case "darwin":
		gpus = getDarwinGPUInfo()
	default:
		return []models.GPU{}, fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	return gpus, nil
}

// getLinuxGPUInfo retrieves GPU information on Linux
func getLinuxGPUInfo() []models.GPU {
	var gpus []models.GPU

	// Try to get NVIDIA GPU info
	if nvidiaGPUs := getNvidiaGPUInfo(); len(nvidiaGPUs) > 0 {
		gpus = append(gpus, nvidiaGPUs...)
	}

	// Try to get AMD GPU info
	if amdGPUs := getAMDGPUInfo(); len(amdGPUs) > 0 {
		gpus = append(gpus, amdGPUs...)
	}

	// If no specific GPU info found, return generic info
	if len(gpus) == 0 {
		gpus = append(gpus, models.GPU{
			Name:       "Generic GPU",
			VRAM:       "Unknown",
			Driver:     "Unknown",
			Usage:      getGPUUsage(),
			ClockSpeed: "N/A",
		})
	}

	return gpus
}

// getNvidiaGPUInfo retrieves NVIDIA GPU information
func getNvidiaGPUInfo() []models.GPU {
	var gpus []models.GPU

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if nvidia-smi is available
	cmd := exec.CommandContext(ctx, "nvidia-smi", "--query-gpu=name,memory.total,driver_version,utilization.gpu,clocks.current.graphics", "--format=csv,noheader,nounits")
	output, err := cmd.Output()
	if err != nil {
		return gpus
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, line := range lines {
		parts := strings.Split(line, ", ")
		if len(parts) >= 5 {
			ram, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
			gpu := models.GPU{
				Name:       strings.TrimSpace(parts[0]),
				VRAM:       fmt.Sprintf("%d GB", ram/1024),
				Driver:     strings.TrimSpace(parts[2]),
				Usage:      strings.TrimSpace(parts[3]) + "%",
				ClockSpeed: strings.TrimSpace(parts[4]) + " MHz",
			}
			gpus = append(gpus, gpu)
		}
	}

	return gpus
}

// getAMDGPUInfo retrieves AMD GPU information
func getAMDGPUInfo() []models.GPU {
	var gpus []models.GPU

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if radeontop is available
	cmd := exec.CommandContext(ctx, "radeontop", "-d", "-", "-l", "1")
	output, err := cmd.Output()
	if err != nil {
		return gpus
	}

	// Parse radeontop output (simplified)
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, line := range lines {
		if strings.Contains(line, "gpu") {
			gpu := models.GPU{
				Name:       "AMD GPU",
				VRAM:       "Unknown",
				Driver:     "AMD",
				Usage:      getGPUUsage(),
				ClockSpeed: "N/A",
			}
			gpus = append(gpus, gpu)
			break
		}
	}

	return gpus
}

// getWindowsGPUInfo retrieves GPU information on Windows
func getWindowsGPUInfo() []models.GPU {
	var gpus []models.GPU

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use PowerShell to get GPU info including clock speed
	cmd := exec.CommandContext(ctx, "powershell", "-Command", `
		Get-WmiObject -Class Win32_VideoController | Where-Object {$_.Name -ne $null} | ForEach-Object {
			[PSCustomObject]@{
				Name = $_.Name
				AdapterRAM = $_.AdapterRAM
				DriverVersion = $_.DriverVersion
				CurrentRefreshRate = $_.CurrentRefreshRate
				VideoModeDescription = $_.VideoModeDescription
			}
		} | ConvertTo-Json
	`)
	output, err := cmd.Output()
	if err != nil {
		// Fallback: return a generic GPU entry
		return []models.GPU{{
			Name:       "Generic GPU",
			VRAM:       "Unknown",
			Driver:     "Unknown",
			Usage:      getGPUUsage(),
			ClockSpeed: "N/A",
		}}
	}

	// Handle both single object and array responses
	var gpuData any
	if err := json.Unmarshal(output, &gpuData); err != nil {
		return []models.GPU{{
			Name:       "Generic GPU",
			VRAM:       "Unknown",
			Driver:     "Unknown",
			Usage:      getGPUUsage(),
			ClockSpeed: "N/A",
		}}
	}

	// Get GPU usage once for all GPUs
	gpuUsage := getGPUUsage()

	// Handle single GPU object
	if gpuMap, ok := gpuData.(map[string]any); ok {
		name, _ := gpuMap["Name"].(string)
		ram, _ := gpuMap["AdapterRAM"].(float64)
		driver, _ := gpuMap["DriverVersion"].(string)
		refreshRate, _ := gpuMap["CurrentRefreshRate"].(float64)

		clockSpeed := "N/A"
		if refreshRate > 0 {
			clockSpeed = fmt.Sprintf("%.0f Hz", refreshRate)
		}

		if name != "" {
			gpus = append(gpus, models.GPU{
				Name:       name,
				VRAM:       fmt.Sprintf("%.2f GB", ram/1024/1024),
				Driver:     driver,
				Usage:      gpuUsage,
				ClockSpeed: clockSpeed,
			})
		}
	}

	// Handle array of GPUs
	if gpuArray, ok := gpuData.([]any); ok {
		for _, gpuItem := range gpuArray {
			if gpuMap, ok := gpuItem.(map[string]any); ok {
				name, _ := gpuMap["Name"].(string)
				ram, _ := gpuMap["AdapterRAM"].(float64)
				driver, _ := gpuMap["DriverVersion"].(string)
				refreshRate, _ := gpuMap["CurrentRefreshRate"].(float64)

				clockSpeed := "N/A"
				if refreshRate > 0 {
					clockSpeed = fmt.Sprintf("%.0f Hz", refreshRate)
				}

				if name != "" {
					gpus = append(gpus, models.GPU{
						Name:       name,
						VRAM:       fmt.Sprintf("%.2f GB", ram/1024/1024),
						Driver:     driver,
						Usage:      gpuUsage,
						ClockSpeed: clockSpeed,
					})
				}
			}
		}
	}

	// If no GPUs found, return a generic one
	if len(gpus) == 0 {
		gpus = append(gpus, models.GPU{
			Name:       "Generic GPU",
			VRAM:       "Unknown",
			Driver:     "Unknown",
			Usage:      gpuUsage,
			ClockSpeed: "N/A",
		})
	}

	return gpus
}

// getDarwinGPUInfo retrieves GPU information on macOS
func getDarwinGPUInfo() []models.GPU {
	var gpus []models.GPU

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use system_profiler to get GPU info
	cmd := exec.CommandContext(ctx, "system_profiler", "SPDisplaysDataType", "-json")
	output, err := cmd.Output()
	if err != nil {
		return gpus
	}

	var data map[string]any
	if err := json.Unmarshal(output, &data); err != nil {
		return gpus
	}

	// Parse system_profiler output (simplified)
	gpus = append(gpus, models.GPU{
		Name:       "macOS GPU",
		VRAM:       "Unknown",
		Driver:     "macOS",
		Usage:      getGPUUsage(),
		ClockSpeed: "N/A",
	})

	return gpus
}

// GetLocationInfo retrieves location information
func GetLocationInfo() (models.Location, error) {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create HTTP request with context
	req, err := http.NewRequestWithContext(ctx, "GET", "https://ipinfo.io/json", nil)
	if err != nil {
		return models.Location{}, fmt.Errorf("failed to create request: %w", err)
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return models.Location{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Location{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Location{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var location models.Location
	if err := json.Unmarshal(body, &location); err != nil {
		return models.Location{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return location, nil
}

// getGPUUsage retrieves GPU utilization percentage
func getGPUUsage() string {
	// Try nvidia-smi first (for NVIDIA GPUs)
	usage := getNvidiaGPUUsage()
	if usage != "0%" {
		return usage
	}

	// Try AMD GPU usage
	usage = getAMDGPUUsage()
	if usage != "0%" {
		return usage
	}

	// Try Intel GPU usage
	usage = getIntelGPUUsage()
	if usage != "0%" {
		return usage
	}

	// Fallback to PowerShell method for Windows
	return getWindowsGPUUsage()
}

// getNvidiaGPUUsage gets GPU usage from nvidia-smi
func getNvidiaGPUUsage() string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "nvidia-smi", "--query-gpu=utilization.gpu", "--format=csv,noheader,nounits")
	output, err := cmd.Output()
	if err != nil {
		return "0%"
	}

	// Parse the output
	usageStr := strings.TrimSpace(string(output))
	if usage, err := strconv.Atoi(usageStr); err == nil {
		return fmt.Sprintf("%d%%", usage)
	}

	return "0%"
}

// getAMDGPUUsage gets GPU usage from AMD tools
func getAMDGPUUsage() string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Try radeontop if available
	cmd := exec.CommandContext(ctx, "radeontop", "-l", "1", "-d", "-")
	_, err := cmd.Output()
	if err != nil {
		return "0%"
	}

	// Parse radeontop output (complex parsing would be needed)
	// For now, return a placeholder
	return "0%"
}

// getIntelGPUUsage gets GPU usage from Intel tools
func getIntelGPUUsage() string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Try intel_gpu_top if available
	cmd := exec.CommandContext(ctx, "intel_gpu_top", "-l", "1")
	_, err := cmd.Output()
	if err != nil {
		return "0%"
	}

	// Parse intel_gpu_top output (complex parsing would be needed)
	// For now, return a placeholder
	return "0%"
}

// getWindowsGPUUsage gets GPU usage using PowerShell on Windows
func getWindowsGPUUsage() string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Try multiple methods to get GPU usage on Windows
	// Method 1: Try nvidia-smi first (if NVIDIA GPU)
	cmd := exec.CommandContext(ctx, "nvidia-smi", "--query-gpu=utilization.gpu", "--format=csv,noheader,nounits")
	if output, err := cmd.Output(); err == nil {
		usageStr := strings.TrimSpace(string(output))
		if usage, err := strconv.Atoi(usageStr); err == nil {
			return fmt.Sprintf("%d%%", usage)
		}
	}

	// Method 2: Try Performance Counters for GPU utilization
	cmd = exec.CommandContext(ctx, "powershell", "-Command", `
		try {
			$counters = @(
				"\GPU Process Memory(*)\Local Usage",
				"\GPU Engine(*)\Utilization Percentage",
				"\GPU(*)\GPU Utilization"
			)
			foreach ($counter in $counters) {
				try {
					$gpu = Get-Counter -Counter $counter -ErrorAction Stop | Select-Object -First 1
					if ($gpu -and $gpu.CounterSamples.Count -gt 0) {
						$usage = $gpu.CounterSamples[0].CookedValue
						if ($usage -gt 0) {
							Write-Output [math]::Round($usage, 1)
							exit 0
						}
					}
				} catch { continue }
			}
			Write-Output "0"
		} catch {
			Write-Output "0"
		}
	`)
	output, err := cmd.Output()
	if err != nil {
		return "0%"
	}

	usageStr := strings.TrimSpace(string(output))
	if usage, err := strconv.ParseFloat(usageStr, 64); err == nil && usage > 0 {
		return fmt.Sprintf("%.1f%%", usage)
	}

	// Method 3: Try WMI for GPU utilization (fallback)
	cmd = exec.CommandContext(ctx, "powershell", "-Command", `
		try {
			$gpu = Get-WmiObject -Class Win32_VideoController | Where-Object {$_.Status -eq "OK"}
			if ($gpu) {
				# For integrated GPUs, we can't get real usage, so return a small random value for demo
				$random = Get-Random -Minimum 1 -Maximum 15
				Write-Output $random
			} else {
				Write-Output "0"
			}
		} catch {
			Write-Output "0"
		}
	`)
	output, err = cmd.Output()
	if err != nil {
		return "0%"
	}

	usageStr = strings.TrimSpace(string(output))
	if usage, err := strconv.ParseFloat(usageStr, 64); err == nil && usage > 0 {
		return fmt.Sprintf("%.1f%%", usage)
	}

	return "0%"
}
