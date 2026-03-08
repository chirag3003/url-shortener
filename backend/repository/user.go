package repository

import (
	"context"
	"errors"
	"strconv"

	"github.com/chirag3003/go-backend-template/db"
	"github.com/chirag3003/go-backend-template/models"
	"github.com/jackc/pgx/v5"
)

//go:generate mockgen -destination=mock/mock_user.go -package=mock github.com/chirag3003/go-backend-template/repository UserRepository

// UserRepository defines the interface for user data access.
type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
}

type userRepository struct {
	conn db.Connection
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(conn db.Connection) UserRepository {
	return &userRepository{conn: conn}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	const q = `
		INSERT INTO users (id, name, email, phone_no, hash, avatar_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())`
	_, err := r.conn.Pool().Exec(ctx, q, user.ID, user.Name, user.Email, nullableString(user.PhoneNo), user.Hash, nullableString(user.AvatarURL))
	return err
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	parsedID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, nil
	}

	const q = `
		SELECT id, name, email, COALESCE(phone_no, ''), hash, COALESCE(avatar_url, ''), created_at, updated_at
		FROM users
		WHERE id = $1`

	var user models.User
	err = r.conn.Pool().QueryRow(ctx, q, parsedID).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PhoneNo,
		&user.Hash,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	const q = `
		SELECT id, name, email, COALESCE(phone_no, ''), hash, COALESCE(avatar_url, ''), created_at, updated_at
		FROM users
		WHERE email = $1`

	var user models.User
	err := r.conn.Pool().QueryRow(ctx, q, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PhoneNo,
		&user.Hash,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	const q = `
		UPDATE users
		SET name = $2, email = $3, phone_no = $4, hash = $5, avatar_url = $6, updated_at = NOW()
		WHERE id = $1`
	_, err := r.conn.Pool().Exec(ctx, q, user.ID, user.Name, user.Email, nullableString(user.PhoneNo), user.Hash, nullableString(user.AvatarURL))
	return err
}
