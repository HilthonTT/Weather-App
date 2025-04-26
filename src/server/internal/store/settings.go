package store

import (
	"context"
	"database/sql"
	"errors"
)

type TempFormat string
type TimeFormat string
type SpeedFormat string

const (
	Celsius    TempFormat = "celsius"
	Fahrenheit TempFormat = "fahrenheit"

	TwentyFourHour TimeFormat = "24h"
	TwelveHour     TimeFormat = "12h"

	Kmph SpeedFormat = "kmph"
	Mph  SpeedFormat = "mph"
)

type Settings struct {
	ID          int64       `json:"id"`
	UserID      int64       `json:"user_id"`
	TempFormat  TempFormat  `json:"temp_format"`
	TimeFormat  TimeFormat  `json:"time_format"`
	SpeedFormat SpeedFormat `json:"speed_format"`
}

func (s *Settings) Update(tempFormat TempFormat, timeFormat TimeFormat, speedFormat SpeedFormat) {
	s.TempFormat = tempFormat
	s.TimeFormat = timeFormat
	s.SpeedFormat = speedFormat
}

type SettingsStore struct {
	db *sql.DB
}

func (s *SettingsStore) Create(ctx context.Context, settings *Settings) error {
	query := `
    INSERT INTO settings (user_id, temp_format, time_format, speed_format)
    VALUES ($1, $2, $3, $4)
    RETURNING id;
  `

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		settings.UserID,
		settings.TempFormat,
		settings.TimeFormat,
		settings.SpeedFormat,
	).Scan(&settings.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *SettingsStore) GetByUserID(ctx context.Context, userID int64) (*Settings, error) {
	query := `
    SELECT id, user_id, temp_format, time_format, speed_format
    FROM settings
    WHERE user_id = $1
  `

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var settings Settings
	err := s.db.QueryRowContext(ctx, query, userID).Scan(
		&settings.ID,
		&settings.UserID,
		&settings.TempFormat,
		&settings.TimeFormat,
		&settings.SpeedFormat,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &settings, nil
}

func (s *SettingsStore) Update(ctx context.Context, settings *Settings) error {
	query := `
		UPDATE settings
		SET temp_format = $1, time_format = $2, speed_format = $3
		WHERE user_id = $4;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	result, err := s.db.ExecContext(
		ctx,
		query,
		settings.TempFormat,
		settings.TimeFormat,
		settings.SpeedFormat,
		settings.UserID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
