package database

import "fmt"

// AddSchedule adds a new schedule to the database
func (db *DB) AddSchedule(schedule *Schedule) error {
	if db == nil || db.conn == nil {
		return fmt.Errorf("database connection not initialized")
	}

	if schedule == nil {
		return fmt.Errorf("schedule cannot be nil")
	}

	query := `
	INSERT INTO schedules (title, start_time, end_time, repeat_pattern, enabled)
	VALUES (?, ?, ?, ?, ?)
	`

	result, err := db.conn.Exec(query, schedule.Title, schedule.StartTime, schedule.EndTime, schedule.RepeatPattern, schedule.Enabled)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	schedule.ID = int(id)
	return nil
}

// GetSchedule retrieves a schedule by ID
func (db *DB) GetSchedule(id int) (*Schedule, error) {
	query := `
	SELECT id, title, start_time, end_time, repeat_pattern, enabled, created_at
	FROM schedules WHERE id = ?
	`

	row := db.conn.QueryRow(query, id)
	schedule := &Schedule{}

	err := row.Scan(
		&schedule.ID,
		&schedule.Title,
		&schedule.StartTime,
		&schedule.EndTime,
		&schedule.RepeatPattern,
		&schedule.Enabled,
		&schedule.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return schedule, nil
}

// GetAllSchedules retrieves all schedules
func (db *DB) GetAllSchedules() ([]*Schedule, error) {
	query := `
	SELECT id, title, start_time, end_time, repeat_pattern, enabled, created_at
	FROM schedules ORDER BY created_at DESC
	`

	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []*Schedule
	for rows.Next() {
		schedule := &Schedule{}
		err := rows.Scan(
			&schedule.ID,
			&schedule.Title,
			&schedule.StartTime,
			&schedule.EndTime,
			&schedule.RepeatPattern,
			&schedule.Enabled,
			&schedule.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	return schedules, nil
}

// UpdateSchedule updates an existing schedule
func (db *DB) UpdateSchedule(schedule *Schedule) error {
	query := `
	UPDATE schedules 
	SET title = ?, start_time = ?, end_time = ?, repeat_pattern = ?, enabled = ?
	WHERE id = ?
	`

	_, err := db.conn.Exec(query, schedule.Title, schedule.StartTime, schedule.EndTime, schedule.RepeatPattern, schedule.Enabled, schedule.ID)
	return err
}

// DeleteSchedule deletes a schedule by ID
func (db *DB) DeleteSchedule(id int) error {
	query := `DELETE FROM schedules WHERE id = ?`
	_, err := db.conn.Exec(query, id)
	return err
}

// ToggleSchedule enables/disables a schedule
func (db *DB) ToggleSchedule(id int, enabled bool) error {
	query := `UPDATE schedules SET enabled = ? WHERE id = ?`
	_, err := db.conn.Exec(query, enabled, id)
	return err
}

// GetEnabledSchedules retrieves all enabled schedules
func (db *DB) GetEnabledSchedules() ([]*Schedule, error) {
	query := `
	SELECT id, title, start_time, end_time, repeat_pattern, enabled, created_at
	FROM schedules WHERE enabled = 1 ORDER BY start_time ASC
	`

	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []*Schedule
	for rows.Next() {
		schedule := &Schedule{}
		err := rows.Scan(
			&schedule.ID,
			&schedule.Title,
			&schedule.StartTime,
			&schedule.EndTime,
			&schedule.RepeatPattern,
			&schedule.Enabled,
			&schedule.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	return schedules, nil
}
