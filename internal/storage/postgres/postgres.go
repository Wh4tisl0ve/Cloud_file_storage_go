package storage

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/domain"
	"github.com/Wh4tisl0ve/Cloud_file_storage_go/pkg/postgres"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepository struct {
	*postgres.Postgres
}

func NewUserRepository(pg *postgres.Postgres) *UserRepository {
	return &UserRepository{pg}
}

func (repo *UserRepository) Save(u *domain.User) error {
	_, err := repo.Conn.Exec(
		"INSERT INTO users(username, password) VALUES ($1, $2)",
		u.Username,
		u.Password,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.ErrUserAlreadyExists
		}
		return fmt.Errorf("failed to insert user record: %w", err)
	}

	return nil
}

func (repo *UserRepository) FindByUsername(username string) (*domain.User, error) {
	var u domain.User

	row := repo.Conn.QueryRow(
		"SELECT id, username, password, created_at FROM users WHERE username = $1",
		username,
	)

	if err := row.Scan(&u.Id, &u.Username, &u.Password, &u.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to fetch user data: %w", err)
	}

	return &u, nil
}
