package app

import (
	"context"
	"log"
	"os"

	"github.com/kishansakhiya/wails-demo/backend/app/services"
)

// App struct
type App struct {
	ctx           context.Context
	systemService *services.SystemService
	logger        *log.Logger
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		systemService: services.NewSystemService(),
		logger:        log.New(os.Stdout, "[SYSTEM-BENCHMARK] ", log.LstdFlags),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.logger.Println("Application started")
}

// shutdown is called at application termination
func (a *App) Shutdown(ctx context.Context) {
	a.logger.Println("Application shutting down")
}

// domReady is called after front-end resources have been loaded
func (a *App) DomReady(ctx context.Context) {
	a.logger.Println("DOM ready")
}

// GetAllSystemInfo retrieves all system information
func (a *App) GetAllSystemInfo() (any, error) {
	return a.systemService.GetAllSystemInfo()
}

// GetCPUInfo retrieves CPU information
func (a *App) GetCPUInfo() (any, error) {
	return a.systemService.GetCPUInfo()
}

// GetGPUInfo retrieves GPU information
func (a *App) GetGPUInfo() (any, error) {
	return a.systemService.GetGPUInfo()
}

// GetOSInfo retrieves OS information
func (a *App) GetOSInfo() (any, error) {
	return a.systemService.GetOSInfo()
}

// GetLocationInfo retrieves location information
func (a *App) GetLocationInfo() (any, error) {
	return a.systemService.GetLocationInfo()
}

// GetMemoryInfo retrieves memory information
func (a *App) GetMemoryInfo() (any, error) {
	return a.systemService.GetMemoryInfo()
}

// GetDiskInfo retrieves disk information
func (a *App) GetDiskInfo() (any, error) {
	return a.systemService.GetDiskInfo()
}

// GetHardwareInfo retrieves hardware information
func (a *App) GetHardwareInfo() (any, error) {
	return a.systemService.GetHardwareInfo()
}

// GetUsagePercentages retrieves usage percentages
func (a *App) GetUsagePercentages() (any, error) {
	return a.systemService.GetUsagePercentages()
}
