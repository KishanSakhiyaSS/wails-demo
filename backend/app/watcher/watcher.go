package watcher

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/kishansakhiya/wails-demo/backend/app/scheduler"
)

// WatcherService handles background synchronization between database and system scheduler
type WatcherService struct {
	schedulerService *scheduler.SchedulerService
	logger           *log.Logger
	interval         time.Duration
}

// NewWatcherService creates a new watcher service
func NewWatcherService(schedulerService *scheduler.SchedulerService, interval time.Duration) *WatcherService {
	return &WatcherService{
		schedulerService: schedulerService,
		logger:           log.New(os.Stdout, "[WATCHER] ", log.LstdFlags),
		interval:         interval,
	}
}

// StartWatcher starts the background watcher service
func (w *WatcherService) StartWatcher(ctx context.Context) {
	w.logger.Printf("Starting watcher service with interval: %v", w.interval)

	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	// Initial sync
	w.logger.Println("Performing initial sync...")
	if err := w.schedulerService.SyncWithSystem(); err != nil {
		w.logger.Printf("Initial sync failed: %v", err)
	}

	for {
		select {
		case <-ticker.C:
			w.logger.Println("Performing scheduled sync...")
			if err := w.schedulerService.SyncWithSystem(); err != nil {
				w.logger.Printf("Sync failed: %v", err)
			} else {
				w.logger.Println("Sync completed successfully")
			}
		case <-ctx.Done():
			w.logger.Println("Watcher service stopping...")
			return
		}
	}
}

// StopWatcher stops the watcher service (called via context cancellation)
func (w *WatcherService) StopWatcher() {
	w.logger.Println("Watcher service stop requested")
}
