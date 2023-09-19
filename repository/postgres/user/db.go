package pguser

import "hangout/repository/postgres"

type DB struct {
	conn *postgres.PgDB
}

func New(coon *postgres.PgDB) *DB {
	return &DB{conn: coon}
}
