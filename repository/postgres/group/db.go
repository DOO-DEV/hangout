package pggroup

import (
	"hangout/repository/postgres"
)

type DB struct {
	conn *postgres.PgDB
}

func New(conn *postgres.PgDB) *DB {
	return &DB{conn: conn}
}
