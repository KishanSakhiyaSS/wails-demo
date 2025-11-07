package scheduler

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/kishansakhiya/wails-demo/backend/app/database"
	"github.com/kishansakhiya/wails-demo/backend/app/utils"
)

const CREATE_NO_WINDOW = 0x08000000

func encodePowerShellCommand(command string) (string, error) {
	utf16Command, err := syscall.UTF16FromString(command)
	if err != nil {
		return "", err
	}
	utf16Command = utf16Command[:len(utf16Command)-1]

	buf := make([]byte, len(utf16Command)*2)
	for i, u := range utf16Command {
		buf[i*2] = byte(u)
		buf[i*2+1] = byte(u >> 8)
	}

	return base64.StdEncoding.EncodeToString(buf), nil
}

// runSchTasks executes schtasks directly with proper argument handling
func runSchTasks(ctx context.Context, args ...string) error {
	cmd := exec.CommandContext(ctx, "schtasks", args...)

	// Initialize SysProcAttr to hide window
	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = &syscall.SysProcAttr{}
	}
	cmd.SysProcAttr.CreationFlags |= CREATE_NO_WINDOW
	cmd.SysProcAttr.HideWindow = true

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("schtasks failed: %w, output: %s", err, string(output))
	}
	return nil
}

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
	executableName = strings.TrimSuffix(executableName, filepath.Ext(executableName))

	// Place tasks under the desired Task Scheduler folder
	baseTaskName := fmt.Sprintf("WailsDemo_Schedule_%d", schedule.ID)
	startTaskName := fmt.Sprintf("%s%s_Start", utils.TaskFolder, baseTaskName)
	endTaskName := fmt.Sprintf("%s%s_End", utils.TaskFolder, baseTaskName)

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

		// Escape the appPath for use in PowerShell command (escape single quotes for PowerShell)
		escapedAppPath := strings.ReplaceAll(appPath, "'", "''")
		startScript := fmt.Sprintf(`Start-Process '%s'`, escapedAppPath)
		encodedStartScript, err := encodePowerShellCommand(startScript)
		if err != nil {
			return fmt.Errorf("failed to encode start task command: %w", err)
		}
		powershellCmd := fmt.Sprintf(`powershell -WindowStyle Hidden -NoProfile -NonInteractive -EncodedCommand %s`, encodedStartScript)

		// Execute schtasks directly with proper arguments
		if err := runSchTasks(context.Background(), "/create", "/tn", startTaskName, "/tr", powershellCmd, "/sc", "once", "/sd", startDate, "/st", startTime, "/f"); err != nil {
			return fmt.Errorf("failed to create start task: %w", err)
		}

		// Escape the executable name for use in PowerShell command
		escapedExecName := strings.ReplaceAll(executableName, "'", "''")
		stopScript := fmt.Sprintf(`Stop-Process -Name '%s' -Force`, escapedExecName)
		encodedStopScript, err := encodePowerShellCommand(stopScript)
		if err != nil {
			runSchTasks(context.Background(), "/delete", "/tn", startTaskName, "/f")
			return fmt.Errorf("failed to encode end task command: %w", err)
		}
		powershellCmd = fmt.Sprintf(`powershell -WindowStyle Hidden -NoProfile -NonInteractive -EncodedCommand %s`, encodedStopScript)

		// Execute schtasks directly with proper arguments
		if err := runSchTasks(context.Background(), "/create", "/tn", endTaskName, "/tr", powershellCmd, "/sc", "once", "/sd", endDate, "/st", endTimeStr, "/f"); err != nil {
			// Clean up start task if end task creation fails
			runSchTasks(context.Background(), "/delete", "/tn", startTaskName, "/f")
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

		// Escape the appPath for use in PowerShell command
		escapedAppPath := strings.ReplaceAll(appPath, "'", "''")
		startScript := fmt.Sprintf(`Start-Process '%s'`, escapedAppPath)
		encodedStartScript, err := encodePowerShellCommand(startScript)
		if err != nil {
			return fmt.Errorf("failed to encode start task command: %w", err)
		}
		powershellCmd := fmt.Sprintf(`powershell -WindowStyle Hidden -NoProfile -NonInteractive -EncodedCommand %s`, encodedStartScript)

		// Execute schtasks directly with proper arguments
		if err := runSchTasks(context.Background(), "/create", "/tn", startTaskName, "/tr", powershellCmd, "/sc", scheduleType, "/st", startTime, "/f"); err != nil {
			return fmt.Errorf("failed to create start task: %w", err)
		}

		// Escape the executable name for use in PowerShell command
		escapedExecName := strings.ReplaceAll(executableName, "'", "''")
		stopScript := fmt.Sprintf(`Stop-Process -Name '%s' -Force`, escapedExecName)
		encodedStopScript, err := encodePowerShellCommand(stopScript)
		if err != nil {
			runSchTasks(context.Background(), "/delete", "/tn", startTaskName, "/f")
			return fmt.Errorf("failed to encode end task command: %w", err)
		}
		powershellCmd = fmt.Sprintf(`powershell -WindowStyle Hidden -NoProfile -NonInteractive -EncodedCommand %s`, encodedStopScript)

		// Execute schtasks directly with proper arguments
		if err := runSchTasks(context.Background(), "/create", "/tn", endTaskName, "/tr", powershellCmd, "/sc", scheduleType, "/st", endTime, "/f"); err != nil {
			// Clean up start task if end task creation fails
			runSchTasks(context.Background(), "/delete", "/tn", startTaskName, "/f")
			return fmt.Errorf("failed to create end task: %w", err)
		}
	}

	s.logger.Printf("Created Windows tasks for schedule %d", schedule.ID)
	return nil
}

func (s *SchedulerService) removeWindowsTasks(schedule *database.Schedule) error {
	baseTaskName := fmt.Sprintf("WailsDemo_Schedule_%d", schedule.ID)
	startTaskName := fmt.Sprintf("%s%s_Start", utils.TaskFolder, baseTaskName)
	endTaskName := fmt.Sprintf("%s%s_End", utils.TaskFolder, baseTaskName)

	// Remove both tasks
	runSchTasks(context.Background(), "/delete", "/tn", startTaskName, "/f")
	runSchTasks(context.Background(), "/delete", "/tn", endTaskName, "/f")

	s.logger.Printf("Removed Windows tasks for schedule %d", schedule.ID)
	return nil
}

func (s *SchedulerService) verifyWindowsTasks(schedule *database.Schedule) error {
	if schedule == nil {
		return fmt.Errorf("schedule cannot be nil")
	}

	baseTaskName := fmt.Sprintf("WailsDemo_Schedule_%d", schedule.ID)
	startTaskName := fmt.Sprintf("%s%s_Start", utils.TaskFolder, baseTaskName)
	endTaskName := fmt.Sprintf("%s%s_End", utils.TaskFolder, baseTaskName)

	// Check if tasks exist
	if err := runSchTasks(context.Background(), "/query", "/tn", startTaskName); err != nil {
		return fmt.Errorf("start task not found")
	}

	if err := runSchTasks(context.Background(), "/query", "/tn", endTaskName); err != nil {
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
