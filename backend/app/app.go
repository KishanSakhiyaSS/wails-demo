package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kishansakhiya/wails-demo/backend/app/database"
	"github.com/kishansakhiya/wails-demo/backend/app/services"
	"github.com/kishansakhiya/wails-demo/backend/app/scheduler"
	"github.com/kishansakhiya/wails-demo/backend/app/watcher"
)

// App struct
type App struct {
	ctx              context.Context
	systemService    *services.SystemService
	schedulerService *scheduler.SchedulerService
	watcherService   *watcher.WatcherService
	db               *database.DB
	logger           *log.Logger
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		systemService: services.NewSystemService(),
		logger:        log.New(os.Stdout, "[APP] ", log.LstdFlags),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.logger.Println("Application started")

	// Initialize database
	var err error
	a.db, err = database.NewDB()
	if err != nil {
		a.logger.Printf("Failed to initialize database: %v", err)
		return
	}
	a.logger.Println("Database initialized successfully")

	// Initialize scheduler service
	a.schedulerService = scheduler.NewSchedulerService(a.db)
	a.logger.Println("Scheduler service initialized")

	// Initialize and start watcher service
	a.watcherService = watcher.NewWatcherService(a.schedulerService, 60*time.Second)
	go a.watcherService.StartWatcher(ctx)
	a.logger.Println("Watcher service started")
}

// shutdown is called at application termination
func (a *App) Shutdown(ctx context.Context) {
	a.logger.Println("Application shutting down")

	// Stop watcher service
	if a.watcherService != nil {
		a.watcherService.StopWatcher()
	}

	// Close database connection
	if a.db != nil {
		if err := a.db.Close(); err != nil {
			a.logger.Printf("Error closing database: %v", err)
		}
	}
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

// Scheduler methods

// AddSchedule adds a new schedule
func (a *App) AddSchedule(schedule *database.Schedule) error {
	if a.schedulerService == nil {
		return fmt.Errorf("scheduler service not initialized")
	}
	return a.schedulerService.AddSchedule(schedule)
}

// ListSchedules retrieves all schedules
func (a *App) ListSchedules() ([]*database.Schedule, error) {
	if a.schedulerService == nil {
		return nil, fmt.Errorf("scheduler service not initialized")
	}
	return a.schedulerService.ListSchedules()
}

// UpdateSchedule updates an existing schedule
func (a *App) UpdateSchedule(schedule *database.Schedule) error {
	if a.schedulerService == nil {
		return fmt.Errorf("scheduler service not initialized")
	}
	return a.schedulerService.UpdateSchedule(schedule)
}

// DeleteSchedule deletes a schedule
func (a *App) DeleteSchedule(id int) error {
	if a.schedulerService == nil {
		return fmt.Errorf("scheduler service not initialized")
	}
	return a.schedulerService.DeleteSchedule(id)
}

// ToggleSchedule enables/disables a schedule
func (a *App) ToggleSchedule(id int, enabled bool) error {
	if a.schedulerService == nil {
		return fmt.Errorf("scheduler service not initialized")
	}
	return a.schedulerService.ToggleSchedule(id, enabled)
}

// SyncWithSystem manually triggers synchronization with system scheduler
func (a *App) SyncWithSystem() error {
	if a.schedulerService == nil {
		return fmt.Errorf("scheduler service not initialized")
	}
	return a.schedulerService.SyncWithSystem()
}

// OnURL handles custom URL scheme requests
// This method is called when the app is opened via a custom URL scheme (e.g., wails-demo://open)
func (a *App) OnURL(url string) {
	a.logger.Printf("Received URL: %s", url)
	
	// Handle different URL paths
	if url == "wails-demo://open" {
		a.logger.Println("Application opened from external source")
		// You can add additional logic here, such as:
		// - Bringing the window to front
		// - Navigating to a specific page
		// - Passing parameters via URL
	}
}
