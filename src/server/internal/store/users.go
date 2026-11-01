package store

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail    = errors.New("a user with that email already exists")
	ErrDuplicateUsername = errors.New("a user with that username already exists")
)

// User is an account. Password is never serialized to JSON.
type User struct {
	ID            int64    `json:"id"`
	Username      string   `json:"username"`
	Email         string   `json:"email"`
	Password      password `json:"-"`
	CreatedAt     string   `json:"created_at"`
	EmailVerified bool     `json:"email_verified"`
	RoleID        int64    `json:"role_id"`
	Role          Role     `json:"role"`
}

type password struct {
	text *string
	hash []byte
}

func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.text = &text
	p.hash = hash

	return nil
}

func (p *password) Compare(text string) error {
	return bcrypt.CompareHashAndPassword(p.hash, []byte(text))
}

type UserStore struct {
	db *sql.DB
}

// Create inserts a user inside tx. Duplicate email/username are reported as
// ErrDuplicateEmail / ErrDuplicateUsername.
func (s *UserStore) Create(ctx context.Context, tx *sql.Tx, user *User) error {
	query := `
		INSERT INTO users (username, password, email, role_id) VALUES 
    ($1, $2, $3, (SELECT id FROM roles WHERE name = $4))
    RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	role := user.Role.Name
	if role == "" {
		role = "user"
	}

	err := tx.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password.hash,
		user.Email,
		role,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" { // unique_violation
			switch pqErr.Constraint {
			case "users_email_key":
				return ErrDuplicateEmail
			case "users_username_key":
				return ErrDuplicateUsername
			}
		}

		return err
	}

	return nil
}

func (s *UserStore) GetByID(ctx context.Context, userID int64) (*User, error) {
	query := `
		SELECT users.id, username, email, password, created_at, email_verified,
		       roles.id, roles.name, roles.level, COALESCE(roles.description, '')
		FROM users
		JOIN roles ON (users.role_id = roles.id)
		WHERE users.id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &User{}
	err := s.db.QueryRowContext(
		ctx,
		query,
		userID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password.hash,
		&user.CreatedAt,
		&user.EmailVerified,
		&user.Role.ID,
		&user.Role.Name,
		&user.Role.Level,
		&user.Role.Description,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	user.RoleID = user.Role.ID

	return user, nil
}

// CreateAndSendEmailVerification creates the user and its email
// verification token in a single transaction. (The email itself is sent by
// the caller.)
func (s *UserStore) CreateAndSendEmailVerification(ctx context.Context, user *User, token string, invitationExp time.Duration) error {
	return withTx(s.db, ctx, func(tx *sql.Tx) error {
		if err := s.Create(ctx, tx, user); err != nil {
			return err
		}

		if err := s.createEmailVerification(ctx, tx, token, invitationExp, user.ID); err != nil {
			return err
		}

		return nil
	})
}

func (s *UserStore) createEmailVerification(ctx context.Context, tx *sql.Tx, token string, exp time.Duration, userID int64) error {
	query := `
    INSERT INTO email_verifications (token, user_id, expiry)
    VALUES ($1, $2, $3)
    `

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, token, userID, time.Now().Add(exp))
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) update(ctx context.Context, tx *sql.Tx, user *User) error {
	query := `
    UPDATE users 
    SET username = $1, 
        email = $2, 
        email_verified = $3 
    WHERE id = $4
    `

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, user.Username, user.Email, user.EmailVerified, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) deleteEmailVerification(ctx context.Context, tx *sql.Tx, userID int64) error {
	query := `DELETE FROM email_verifications WHERE user_id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) delete(ctx context.Context, tx *sql.Tx, id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a user and any pending email verification.
func (s *UserStore) Delete(ctx context.Context, userID int64) error {
	return withTx(s.db, ctx, func(tx *sql.Tx) error {
		if err := s.deleteEmailVerification(ctx, tx, userID); err != nil {
			return err
		}

		if err := s.delete(ctx, tx, userID); err != nil {
			return err
		}

		return nil
	})
}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, username, email, password, created_at 
    FROM users
		WHERE email = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &User{}
	err := s.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password.hash,
		&user.CreatedAt,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}

func (s *UserStore) getUserFromInvitation(ctx context.Context, tx *sql.Tx, token string) (*User, error) {
	query := `
		SELECT u.id, u.username, u.email, u.created_at, u.email_verified
		FROM users u
		JOIN email_verifications ui ON u.id = ui.user_id
		WHERE ui.token = $1 AND ui.expiry > $2
	`

	hash := sha256.Sum256([]byte(token))
	hashToken := hex.EncodeToString(hash[:])

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &User{}
	err := tx.QueryRowContext(ctx, query, hashToken, time.Now()).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.EmailVerified,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}

// Verify marks the user owning the (plain-text) token as verified and
// consumes the token. Returns ErrNotFound for unknown or expired tokens.
func (s *UserStore) Verify(ctx context.Context, token string) error {
	return withTx(s.db, ctx, func(tx *sql.Tx) error {

		user, err := s.getUserFromInvitation(ctx, tx, token)
		if err != nil {
			return err
		}

		user.EmailVerified = true
		if err := s.update(ctx, tx, user); err != nil {
			return err
		}

		if err := s.deleteEmailVerification(ctx, tx, user.ID); err != nil {
			return err
		}

		return nil
	})
}
