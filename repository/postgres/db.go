package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"time"
)

type Config struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Host     string `koanf:"host"`
	Name     string `koanf:"name"`
}

type PgDB struct {
	config Config
	db     *sql.DB
}

func (d *PgDB) Conn() *sql.DB {
	return d.db
}

func New(cfg Config) *PgDB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Name)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &PgDB{config: cfg, db: db}
}

func (d *PgDB) IsDuplicateKeyError(err error) bool {
	var pgErr *pq.Error
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return true
	}

	return false
}

func (d *PgDB) IsEmptyRowError(err error) bool {
	if errors.Is(err, sql.ErrNoRows) {
		return true
	}

	return false
}

func (d *PgDB) IsForeignKeyError(err error) bool {
	var pgErr *pq.Error
	if errors.As(err, &pgErr) && pgErr.Code == "23503" {
		return true
	}

	return false
}
