package postgres

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Postgres struct {
	dsn  string
	Conn *sql.DB
}

func New(dsn string) (*Postgres, error) {
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return &Postgres{
		dsn:  dsn,
		Conn: conn,
	}, nil
}

func (p *Postgres) Close() error {
	err := p.Conn.Close()

	if err != nil {
		return err
	}

	return nil
}
