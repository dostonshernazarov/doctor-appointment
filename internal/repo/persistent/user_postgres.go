package persistent

import (
	"context"
	"fmt"
	"time"

	"github.com/dostonshernazarov/doctor-appointment/internal/entity"
	"github.com/dostonshernazarov/doctor-appointment/pkg/postgres"
)

const _defaultEntityCap = 64

// UserRepo -.
type UserRepo struct {
	*postgres.Postgres
}

// NewUser -.
func NewUser(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

// Create -.
func (r *UserRepo) CreateUser(ctx context.Context, user entity.User) error {
	sql, args, err := r.Builder.
		Insert("users").
		Columns("fullname", "email", "phone", "password_hash", "role").
		Values(user.FullName, user.Email, user.Phone, user.Password, user.Role).ToSql()

	if err != nil {
		return fmt.Errorf("UserRepo - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserRepo - Store - r.Pool.Exec: %w", err)
	}

	return nil
}

// GetUserByEmail -.
func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	sql, args, err := r.Builder.
		Select("id", "fullname", "email", "phone", "role", "created_at", "updated_at").
		From("users").
		Where("email = ?", email).
		Limit(1).
		ToSql()

	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - GetByEmail - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	var user entity.User
	err = row.Scan(&user.ID, &user.FullName, &user.Email, &user.Phone, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - GetByEmail - row.Scan: %w", err)
	}

	return user, nil
}

// ListUsers -.
func (r *UserRepo) ListUsers(ctx context.Context) ([]entity.User, error) {
	sql, args, err := r.Builder.
		Select("id", "fullname", "email", "phone", "role", "created_at", "updated_at").
		From("users").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("UserRepo - ListUser - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("UserRepo - ListUser - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	users := make([]entity.User, 0, _defaultEntityCap)
	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.ID, &user.FullName, &user.Email, &user.Phone, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("UserRepo - ListUser - rows.Scan: %w", err)
		}

		users = append(users, user)
	}

	return users, nil
}

// UpdateUser -.
func (r *UserRepo) UpdateUser(ctx context.Context, user entity.User) error {
	updateTime := time.Now()
	sql, args, err := r.Builder.
		Update("users").
		Set("fullname", user.FullName).
		Set("email", user.Email).
		Set("phone", user.Phone).
		Set("password", user.Password).
		Set("updated_at", updateTime).
		Where("id = ?", user.ID).
		ToSql()

	if err != nil {
		return fmt.Errorf("UserRepo - UpdateUser - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserRepo - UpdateUser - r.Pool.Exec: %w", err)
	}

	return nil
}

// DeleteUser -.
func (r *UserRepo) DeleteUser(ctx context.Context, id int) error {
	sql, args, err := r.Builder.
		Delete("users").
		Where("id = ?", id).
		ToSql()

	if err != nil {
		return fmt.Errorf("UserRepo - DeleteUser - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserRepo - DeleteUser - r.Pool.Exec: %w", err)
	}

	return nil
}

// GetUserByID -.
func (r *UserRepo) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	sql, args, err := r.Builder.
		Select("id", "fullname", "email", "phone", "role", "created_at", "updated_at").
		From("users").
		Where("id = ?", id).
		Limit(1).
		ToSql()

	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - GetUserByID - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	var user entity.User
	err = row.Scan(&user.ID, &user.FullName, &user.Email, &user.Phone, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - GetUserByID - row.Scan: %w", err)
	}

	return user, nil

}

// GetPasswordHash -.
func (r *UserRepo) GetPasswordHash(ctx context.Context, email string) (string, error) {
	sql, args, err := r.Builder.
		Select("password_hash").
		From("users").
		Where("email = ?", email).
		Limit(1).
		ToSql()

	if err != nil {
		return "", fmt.Errorf("UserRepo - GetPasswordHash - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	var passwordHash string
	err = row.Scan(&passwordHash)
	if err != nil {
		return "", fmt.Errorf("UserRepo - GetPasswordHash - row.Scan: %w", err)
	}

	return passwordHash, nil
}

// UpdateToken -.
func (r *UserRepo) UpdateToken(ctx context.Context, id int, token string) error {
	sql, args, err := r.Builder.
		Update("users").
		Set("token", token).
		Where("id = ?", id).
		ToSql()

	if err != nil {
		return fmt.Errorf("UserRepo - UpdateToken - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserRepo - UpdateToken - r.Pool.Exec: %w", err)
	}

	return nil
}
