package store

import (
	"context"
	"database/sql"
	"sync"
	"time"
)

// NewMockStore returns an in-memory Storage for handler tests.
func NewMockStore() Storage {
	return Storage{
		Users:    &MockUserStore{},
		Settings: NewMockSettingsStore(),
	}
}

type MockUserStore struct{}

func (m *MockUserStore) Create(ctx context.Context, tx *sql.Tx, u *User) error {
	return nil
}

func (m *MockUserStore) GetByID(ctx context.Context, userID int64) (*User, error) {
	return &User{ID: userID}, nil
}

func (m *MockUserStore) GetByEmail(context.Context, string) (*User, error) {
	return &User{}, nil
}

func (m *MockUserStore) Delete(ctx context.Context, id int64) error {
	return nil
}

func (m *MockUserStore) CreateAndSendEmailVerification(ctx context.Context, user *User, token string, exp time.Duration) error {
	return nil
}

func (m *MockUserStore) Verify(ctx context.Context, t string) error {
	return nil
}

// MockSettingsStore keeps settings in memory, keyed by user ID.
type MockSettingsStore struct {
	mu   sync.Mutex
	data map[int64]Settings
}

func NewMockSettingsStore() *MockSettingsStore {
	return &MockSettingsStore{data: map[int64]Settings{}}
}

func (m *MockSettingsStore) Create(_ context.Context, s *Settings) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	s.ID = int64(len(m.data) + 1)
	m.data[s.UserID] = *s
	return nil
}

func (m *MockSettingsStore) GetByUserID(_ context.Context, userID int64) (*Settings, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	s, ok := m.data[userID]
	if !ok {
		return nil, ErrNotFound
	}
	return &s, nil
}

func (m *MockSettingsStore) Update(_ context.Context, s *Settings) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[s.UserID]; !ok {
		return ErrNotFound
	}
	m.data[s.UserID] = *s
	return nil
}
