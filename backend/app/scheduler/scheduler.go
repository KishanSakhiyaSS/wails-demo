package scheduler

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/kishansakhiya/wails-demo/backend/app/database"
)

// SchedulerService handles system-level scheduling operations
type SchedulerService struct {
	db     *database.DB
	logger *log.Logger
}

// NewSchedulerService creates a new scheduler service
func NewSchedulerService(db *database.DB) *SchedulerService {
	return &SchedulerService{
		db:     db,
		logger: log.New(os.Stdout, "[SCHEDULER] ", log.LstdFlags),
	}
}

// AddSchedule adds a new schedule and creates system tasks
func (s *SchedulerService) AddSchedule(schedule *database.Schedule) error {
	if schedule == nil {
		return fmt.Errorf("schedule cannot be nil")
	}

	if s.db == nil {
		return fmt.Errorf("database not initialized")
	}

	s.logger.Printf("Adding schedule: %s", schedule.Title)

	// Add to database
	if err := s.db.AddSchedule(schedule); err != nil {
		return fmt.Errorf("failed to add schedule to database: %w", err)
	}

	// Create system tasks if enabled
	if schedule.Enabled {
		if err := s.createSystemTasks(schedule); err != nil {
			s.logger.Printf("Warning: Failed to create system tasks for schedule %d: %v", schedule.ID, err)
			// Don't fail the operation, just log the warning
		}
	}

	return nil
}

// ListSchedules retrieves all schedules
func (s *SchedulerService) ListSchedules() ([]*database.Schedule, error) {
	return s.db.GetAllSchedules()
}

// DeleteSchedule deletes a schedule and removes system tasks
func (s *SchedulerService) DeleteSchedule(id int) error {
	s.logger.Printf("Deleting schedule ID: %d", id)

	// Get schedule first to remove system tasks
	schedule, err := s.db.GetSchedule(id)
	if err != nil {
		return fmt.Errorf("failed to get schedule: %w", err)
	}

	// Remove system tasks
	if err := s.removeSystemTasks(schedule); err != nil {
		s.logger.Printf("Warning: Failed to remove system tasks for schedule %d: %v", id, err)
	}

	// Delete from database
	return s.db.DeleteSchedule(id)
}

// UpdateSchedule updates a schedule and recreates system tasks
func (s *SchedulerService) UpdateSchedule(schedule *database.Schedule) error {
	s.logger.Printf("Updating schedule ID: %d", schedule.ID)

	// Get old schedule to remove old system tasks
	oldSchedule, err := s.db.GetSchedule(schedule.ID)
	if err != nil {
		return fmt.Errorf("failed to get old schedule: %w", err)
	}

	// Remove old system tasks
	if err := s.removeSystemTasks(oldSchedule); err != nil {
		s.logger.Printf("Warning: Failed to remove old system tasks: %v", err)
	}

	// Update in database
	if err := s.db.UpdateSchedule(schedule); err != nil {
		return fmt.Errorf("failed to update schedule in database: %w", err)
	}

	// Create new system tasks if enabled
	if schedule.Enabled {
		if err := s.createSystemTasks(schedule); err != nil {
			s.logger.Printf("Warning: Failed to create new system tasks: %v", err)
		}
	}

	return nil
}

// ToggleSchedule enables/disables a schedule
func (s *SchedulerService) ToggleSchedule(id int, enabled bool) error {
	s.logger.Printf("Toggling schedule ID: %d to enabled: %t", id, enabled)

	// Get schedule
	schedule, err := s.db.GetSchedule(id)
	if err != nil {
		return fmt.Errorf("failed to get schedule: %w", err)
	}

	// Update enabled status
	schedule.Enabled = enabled
	if err := s.db.UpdateSchedule(schedule); err != nil {
		return fmt.Errorf("failed to update schedule: %w", err)
	}

	// Handle system tasks
	if enabled {
		// Create system tasks
		if err := s.createSystemTasks(schedule); err != nil {
			s.logger.Printf("Warning: Failed to create system tasks: %v", err)
		}
	} else {
		// Remove system tasks
		if err := s.removeSystemTasks(schedule); err != nil {
			s.logger.Printf("Warning: Failed to remove system tasks: %v", err)
		}
	}

	return nil
}

// SyncWithSystem ensures database and system scheduler are in sync
func (s *SchedulerService) SyncWithSystem() error {
	s.logger.Println("Syncing with system scheduler")

	// Get all enabled schedules from database
	schedules, err := s.db.GetEnabledSchedules()
	if err != nil {
		return fmt.Errorf("failed to get enabled schedules: %w", err)
	}

	// Check each schedule's system tasks
	for _, schedule := range schedules {
		if err := s.verifySystemTasks(schedule); err != nil {
			s.logger.Printf("Schedule %d system tasks out of sync: %v", schedule.ID, err)
			// Recreate system tasks
			if err := s.createSystemTasks(schedule); err != nil {
				s.logger.Printf("Failed to recreate system tasks for schedule %d: %v", schedule.ID, err)
			}
		}
	}

	return nil
}

// createSystemTasks creates system-level tasks for a schedule
func (s *SchedulerService) createSystemTasks(schedule *database.Schedule) error {
	switch runtime.GOOS {
	case "windows":
		return s.createWindowsTasks(schedule)
	case "linux", "darwin":
		return s.createUnixTasks(schedule)
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

// removeSystemTasks removes system-level tasks for a schedule
func (s *SchedulerService) removeSystemTasks(schedule *database.Schedule) error {
	switch runtime.GOOS {
	case "windows":
		return s.removeWindowsTasks(schedule)
	case "linux", "darwin":
		return s.removeUnixTasks(schedule)
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

// verifySystemTasks checks if system tasks exist for a schedule
func (s *SchedulerService) verifySystemTasks(schedule *database.Schedule) error {
	switch runtime.GOOS {
	case "windows":
		return s.verifyWindowsTasks(schedule)
	case "linux", "darwin":
		return s.verifyUnixTasks(schedule)
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}
