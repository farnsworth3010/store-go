package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	blogTable  = "blog"
	usersTable = "users"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CountRows(table string, db *sqlx.DB) int {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)
	var count int
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0
	}
	return count
}
