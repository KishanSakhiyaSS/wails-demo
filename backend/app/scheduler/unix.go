package scheduler

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/kishansakhiya/wails-demo/backend/app/database"
)

// Unix task management (Linux/macOS)
func (s *SchedulerService) createUnixTasks(schedule *database.Schedule) error {
	appPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}

	// Create cron entries
	startCronEntry := s.buildCronEntry(schedule.StartTime, schedule.RepeatPattern, appPath, "start", schedule.ID)
	endCronEntry := s.buildCronEntry(schedule.EndTime, schedule.RepeatPattern, fmt.Sprintf("pkill -f %s", filepath.Base(appPath)), "end", schedule.ID)

	// Add to crontab
	if err := s.addCronEntry(startCronEntry); err != nil {
		return fmt.Errorf("failed to add start cron entry: %w", err)
	}

	if err := s.addCronEntry(endCronEntry); err != nil {
		// Clean up start entry if end entry fails
		s.removeCronEntry(schedule.ID, "start")
		return fmt.Errorf("failed to add end cron entry: %w", err)
	}

	s.logger.Printf("Created Unix cron entries for schedule %d", schedule.ID)
	return nil
}

func (s *SchedulerService) removeUnixTasks(schedule *database.Schedule) error {
	// Remove cron entries
	s.removeCronEntry(schedule.ID, "start")
	s.removeCronEntry(schedule.ID, "end")

	s.logger.Printf("Removed Unix cron entries for schedule %d", schedule.ID)
	return nil
}

func (s *SchedulerService) verifyUnixTasks(schedule *database.Schedule) error {
	// Check if cron entries exist
	cmd := exec.Command("crontab", "-l")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to list crontab: %w", err)
	}

	crontabContent := string(output)
	scheduleID := strconv.Itoa(schedule.ID)

	// Check for both start and end entries
	if !strings.Contains(crontabContent, fmt.Sprintf("WailsDemo_Schedule_%s_start", scheduleID)) ||
		!strings.Contains(crontabContent, fmt.Sprintf("WailsDemo_Schedule_%s_end", scheduleID)) {
		return fmt.Errorf("cron entries not found")
	}

	return nil
}

func (s *SchedulerService) buildCronEntry(scheduleTime time.Time, repeatPattern, command, taskType string, scheduleID int) string {
	minute := scheduleTime.Minute()
	hour := scheduleTime.Hour()
	dayOfWeek := "*"
	dayOfMonth := "*"
	month := "*"

	switch repeatPattern {
	case "daily":
		// Every day at the specified time
		return fmt.Sprintf("%d %d * * * %s # WailsDemo_Schedule_%d_%s", minute, hour, command, scheduleID, taskType)
	case "weekly":
		// Every week on the same day
		dayOfWeek = strconv.Itoa(int(scheduleTime.Weekday()))
		return fmt.Sprintf("%d %d * * %s %s # WailsDemo_Schedule_%d_%s", minute, hour, dayOfWeek, command, scheduleID, taskType)
	default: // once
		// Specific date and time
		dayOfMonth = strconv.Itoa(scheduleTime.Day())
		month = strconv.Itoa(int(scheduleTime.Month()))
		return fmt.Sprintf("%d %d %s %s * %s # WailsDemo_Schedule_%d_%s", minute, hour, dayOfMonth, month, command, scheduleID, taskType)
	}
}

func (s *SchedulerService) addCronEntry(cronEntry string) error {
	// Get current crontab
	cmd := exec.Command("crontab", "-l")
	currentCrontab, err := cmd.Output()
	if err != nil && !strings.Contains(err.Error(), "no crontab") {
		return err
	}

	// Add new entry
	newCrontab := string(currentCrontab)
	if newCrontab != "" && !strings.HasSuffix(newCrontab, "\n") {
		newCrontab += "\n"
	}
	newCrontab += cronEntry + "\n"

	// Write new crontab
	cmd = exec.Command("crontab", "-")
	cmd.Stdin = strings.NewReader(newCrontab)
	return cmd.Run()
}

func (s *SchedulerService) removeCronEntry(scheduleID int, taskType string) error {
	// Get current crontab
	cmd := exec.Command("crontab", "-l")
	currentCrontab, err := cmd.Output()
	if err != nil {
		return err
	}

	// Remove entries for this schedule
	lines := strings.Split(string(currentCrontab), "\n")
	var newLines []string
	pattern := fmt.Sprintf("WailsDemo_Schedule_%d_%s", scheduleID, taskType)

	for _, line := range lines {
		if !strings.Contains(line, pattern) {
			newLines = append(newLines, line)
		}
	}

	// Write updated crontab
	if len(newLines) > 0 {
		newCrontab := strings.Join(newLines, "\n")
		if !strings.HasSuffix(newCrontab, "\n") {
			newCrontab += "\n"
		}
		cmd = exec.Command("crontab", "-")
		cmd.Stdin = strings.NewReader(newCrontab)
		return cmd.Run()
	} else {
		// Remove entire crontab if no entries left
		cmd = exec.Command("crontab", "-r")
		return cmd.Run()
	}
}


