package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	db *sql.DB
}

func New(dsn string) (*Repository, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}
