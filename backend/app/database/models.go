package database

import "time"

// Schedule represents a scheduled task
type Schedule struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	RepeatPattern string    `json:"repeat_pattern"` // once, daily, weekly
	Enabled       bool      `json:"enabled"`
	CreatedAt     time.Time `json:"created_at"`
}

