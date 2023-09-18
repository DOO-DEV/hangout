package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"hangout/repository/postgres"
	"log"
	"os"
)

type Migrator struct {
	dialect    string
	dbConfig   postgres.Config
	migrations *migrate.FileMigrationSource
}

func (m Migrator) Up() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Name)
	db, err := sql.Open(m.dialect, dsn)
	if err != nil {
		log.Fatalf("can't open postgres: %s", err)
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		log.Fatalf("can't apply migrations: %s", err)
	}

	log.Printf("%d migrations applied", n)
}

func (m Migrator) Down() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Name)
	db, err := sql.Open(m.dialect, dsn)
	if err != nil {
		log.Fatalf("can't open postgres: %s", err)
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Down)
	if err != nil {
		log.Fatalf("can't apply migrations: %s", err)
	}

	log.Printf("%d migrations applied", n)

}

func main() {
	m := Migrator{
		dialect: "postgres",
		dbConfig: postgres.Config{
			Username: "doo-dev",
			Password: "123456",
			Host:     "localhost",
			Name:     "hangout",
		},
		migrations: &migrate.FileMigrationSource{
			Dir: "repository/postgres/migrations",
		},
	}

	cmd := os.Args[1]
	switch cmd {
	case "migrate-up":
		m.Up()
	case "migrate-down":
		m.Down()
	default:
		fmt.Println("wrong option")
	}
}
