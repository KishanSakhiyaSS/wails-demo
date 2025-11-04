package scheduler

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/kishansakhiya/wails-demo/backend/app/database"
	"github.com/kishansakhiya/wails-demo/backend/app/utils"
)

// Windows task management
func (s *SchedulerService) createWindowsTasks(schedule *database.Schedule) error {
	if schedule == nil {
		return fmt.Errorf("schedule cannot be nil")
	}

	if s.logger == nil {
		return fmt.Errorf("logger not initialized")
	}

	appPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}

	// Get the executable name for taskkill
	executableName := filepath.Base(appPath)

	taskName := fmt.Sprintf("WailsDemo_Schedule_%d", schedule.ID)
	startTaskName := fmt.Sprintf("%s_Start", taskName)
	endTaskName := fmt.Sprintf("%s_End", taskName)

	// Create tasks based on repeat pattern
	if schedule.RepeatPattern == "once" {
		// For one-time tasks, use specific date and time
		now := time.Now()
		startTimeValue := schedule.StartTime.Local()
		endTimeValue := schedule.EndTime.Local()

		// Windows Task Scheduler cannot schedule tasks in the past
		if schedule.StartTime.Before(now) {
			daysToAdd := 1
			for schedule.StartTime.AddDate(0, 0, daysToAdd).Before(now.Add(time.Minute)) {
				daysToAdd++
			}
			startTimeValue = schedule.StartTime.AddDate(0, 0, daysToAdd)
			s.logger.Printf("Start time %s is in the past (now: %s), adjusting to: %s",
				schedule.StartTime, now, startTimeValue)
		}

		if schedule.EndTime.Before(now) {
			daysToAdd := 1
			for schedule.EndTime.AddDate(0, 0, daysToAdd).Before(now.Add(time.Minute)) {
				daysToAdd++
			}
			endTimeValue = schedule.EndTime.AddDate(0, 0, daysToAdd)
			s.logger.Printf("End time %s is in the past (now: %s), adjusting to: %s",
				schedule.EndTime, now, endTimeValue)
		}

		// Use schtasks for one-time tasks (requires MM/dd/yyyy for /sd and HH:mm for /st)
		startDate := startTimeValue.Local().Format("02/01/2006")
		startTime := startTimeValue.Local().Format("15:04")
		endDate := endTimeValue.Local().Format("02/01/2006")
		endTimeStr := endTimeValue.Local().Format("15:04")

		startCmd := fmt.Sprintf(`schtasks /create /tn "%s" /tr "%s" /sc once /sd %s /st %s /f`,
			startTaskName, appPath, startDate, startTime)
		if err := utils.RunCommand(context.Background(), startCmd); err != nil {
			return fmt.Errorf("failed to create start task: %w", err)
		}

		endCmd := fmt.Sprintf(`schtasks /create /tn "%s" /tr "cmd /c taskkill /IM %s /F" /sc once /sd %s /st %s /f`,
			endTaskName, executableName, endDate, endTimeStr)
		if err := utils.RunCommand(context.Background(), endCmd); err != nil {
			// Clean up start task if end task creation fails
			utils.RunCommand(context.Background(), fmt.Sprintf(`schtasks /delete /tn "%s" /f`, startTaskName))
			return fmt.Errorf("failed to create end task: %w", err)
		}
	} else {
		// For recurring tasks (daily/weekly)
		scheduleType := s.getWindowsScheduleType(schedule.RepeatPattern)

		// Convert to local time and format for schtasks (HH:mm format)
		startTimeValue := schedule.StartTime.Local()
		endTimeValue := schedule.EndTime.Local()
		startTime := startTimeValue.Format("15:04")
		endTime := endTimeValue.Format("15:04")

		startCmd := fmt.Sprintf(`schtasks /create /tn "%s" /tr "%s" /sc %s /st %s /f`,
			startTaskName, appPath, scheduleType, startTime)

		if err := utils.RunCommand(context.Background(), startCmd); err != nil {
			return fmt.Errorf("failed to create start task: %w", err)
		}

		endCmd := fmt.Sprintf(`schtasks /create /tn "%s" /tr "cmd /c taskkill /IM %s /F" /sc %s /st %s /f`,
			endTaskName, executableName, scheduleType, endTime)

		if err := utils.RunCommand(context.Background(), endCmd); err != nil {
			// Clean up start task if end task creation fails
			utils.RunCommand(context.Background(), fmt.Sprintf(`schtasks /delete /tn "%s" /f`, startTaskName))
			return fmt.Errorf("failed to create end task: %w", err)
		}
	}

	s.logger.Printf("Created Windows tasks for schedule %d", schedule.ID)
	return nil
}

func (s *SchedulerService) removeWindowsTasks(schedule *database.Schedule) error {
	taskName := fmt.Sprintf("WailsDemo_Schedule_%d", schedule.ID)
	startTaskName := fmt.Sprintf("%s_Start", taskName)
	endTaskName := fmt.Sprintf("%s_End", taskName)

	// Remove both tasks
	utils.RunCommand(context.Background(), fmt.Sprintf(`schtasks /delete /tn "%s" /f`, startTaskName))
	utils.RunCommand(context.Background(), fmt.Sprintf(`schtasks /delete /tn "%s" /f`, endTaskName))

	s.logger.Printf("Removed Windows tasks for schedule %d", schedule.ID)
	return nil
}

func (s *SchedulerService) verifyWindowsTasks(schedule *database.Schedule) error {
	if schedule == nil {
		return fmt.Errorf("schedule cannot be nil")
	}

	taskName := fmt.Sprintf("WailsDemo_Schedule_%d", schedule.ID)
	startTaskName := fmt.Sprintf("%s_Start", taskName)
	endTaskName := fmt.Sprintf("%s_End", taskName)

	// Check if tasks exist
	cmd := fmt.Sprintf(`schtasks /query /tn "%s"`, startTaskName)
	if err := utils.RunCommand(context.Background(), cmd); err != nil {
		return fmt.Errorf("start task not found")
	}

	cmd = fmt.Sprintf(`schtasks /query /tn "%s"`, endTaskName)
	if err := utils.RunCommand(context.Background(), cmd); err != nil {
		return fmt.Errorf("end task not found")
	}

	return nil
}

func (s *SchedulerService) getWindowsScheduleType(repeatPattern string) string {
	switch repeatPattern {
	case "daily":
		return "daily"
	case "weekly":
		return "weekly"
	default:
		return "once"
	}
}
