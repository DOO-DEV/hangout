package pgmessage

import "hangout/repository/postgres"

type DB struct {
	conn *postgres.PgDB
}

func New(pgDb *postgres.PgDB) DB {
	return DB{conn: pgDb}
}
