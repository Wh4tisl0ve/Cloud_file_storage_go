package repository

import (
	"database/sql"
	"fmt"

	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/entity"
	"github.com/Wh4tisl0ve/Cloud_file_storage_go/pkg/postgres"
)

type UserRepository struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *UserRepository {
	return &UserRepository{
		pg,
	}
}

func (repo *UserRepository) CreateUser(u *entity.User) error {
	_, err := repo.Conn.Exec(
		"INSERT INTO users(username, password) VALUES ($1, $2)",
		u.Username, u.Password,
	)
	if err != nil {
		return fmt.Errorf("Ошибка добавления новой записи: %s", err.Error())
	}

	return nil
}

func (repo *UserRepository) FindByUsername(userName string) (entity.User, error) {
	var u entity.User

	row := repo.Conn.QueryRow("SELECT * FROM users WHERE username = $1", userName)
	if err := row.Scan(&u.Id, &u.Username, &u.Password, &u.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return u, fmt.Errorf("Нет записей с username = %s", userName)
		}
		return u, fmt.Errorf("Ошибка извлечения данных: %s", err)
	}

	return u, nil
}
