package settings

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hilthontt/weather/internal/db"
)

var (
	ErrDuplicateUserID = errors.New("the user already has a settings")
)

type SettingsStore struct {
	db *sql.DB
}

func NewSettingsStore(db *sql.DB) *SettingsStore {
	return &SettingsStore{db}
}

func (s *SettingsStore) Create(ctx context.Context, settings Settings) error {
	query := `
    INSERT INTO settings (user_id, temperature, wind_speed, time_display)
    VALUES ($1, $2, $3, $4)
    RETRUNING id;
  `

	ctx, cancel := context.WithTimeout(ctx, db.QueryTimeoutDuration)
	defer cancel()

	var id int64

	err := s.db.QueryRowContext(
		ctx,
		query,
		settings.UserID,
		settings.Temperature,
		settings.WindSpeed,
		settings.TimeDisplay,
	).Scan(&id)

	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "settings_user_id_key"`:
			return ErrDuplicateUserID
		}

		return err
	}

	return nil
}

func (s *SettingsStore) GetByUserID(ctx context.Context, userID int64) (*Settings, error) {
	query := `
    SELECT id, user_id, temperature, wind_speed, time_display
    FROM settings
    WHERE id = $1
  `

	ctx, cancel := context.WithTimeout(ctx, db.QueryTimeoutDuration)
	defer cancel()

	settings := &Settings{}
	err := s.db.QueryRowContext(ctx, query, userID).Scan(
		&settings.ID,
		&settings.Temperature,
		&settings.WindSpeed,
		&settings.TimeDisplay,
	)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, db.ErrNotFound
		default:
			return nil, err
		}
	}

	return settings, nil
}

func (s *SettingsStore) Update(
	ctx context.Context,
	userID int64,
	temperature TemperatureUnit,
	windSpeed WindSpeedUnit,
	timeDisplay TimeFormat,
) error {
	query := `
		UPDATE settings
		SET temperature = $1,
		    wind_speed = $2,
		    time_display = $3
		WHERE user_id = $4;
	`
	ctx, cancel := context.WithTimeout(ctx, db.QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, temperature, windSpeed, timeDisplay, userID)
	if err != nil {
		return err
	}

	return nil
}
