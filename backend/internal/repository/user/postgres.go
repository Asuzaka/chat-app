package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserPostgresRepository struct {
	db *pgxpool.Pool
}

func NewUserPostgresRepository(db *pgxpool.Pool) *UserPostgresRepository {
	return &UserPostgresRepository{db: db}
}

func (r *UserPostgresRepository) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (id, username, email, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING created_at, updated_at
	`

	return r.db.QueryRow(ctx, query,
		user.ID,
		user.Username,
		user.Email,
		user.PasswordHash,
	).Scan(&user.CreatedAt, &user.UpdatedAt)
}

func (r *UserPostgresRepository) FindByID(ctx context.Context, id string) (*User, error) {
	query := `
	 	SELECT id, username, email, password_hash, avatar_url, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user User

	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserPostgresRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, username, email, password_hash, avatar_url, created_at, updated_at
		FROM users
		WHERE email = $1
	`
	var user User

	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserPostgresRepository) FindByUsername(ctx context.Context, username string) (*User, error) {

	query := `
		SELECT id, username, email, password_hash, avatar_url, created_at, updated_at
		FROM users
		WHERE username = $1 `

	var user User

	err := r.db.QueryRow(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserPostgresRepository) Update(ctx context.Context, user *User) error {
	query := `
		UPDATE users
		SET username = $1, email = $2, password_hash = $3, avatar_url = $4, updated_at = now()
		Where id = $5
	`

	cmdTag, err := r.db.Exec(ctx, query,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.AvatarURL,
		user.ID)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return NoUserFoundError
	}

	return nil
}

func (r *UserPostgresRepository) Delete(ctx context.Context, id string) error {
	query := `
		Delete FROM users
		WHERE id = $1
	`

	cmdTag, err := r.db.Exec(ctx, query, id)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return NoUserFoundError
	}

	return nil
}
