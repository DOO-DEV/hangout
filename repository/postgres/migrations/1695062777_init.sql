-- +migrate Up
CREATE TABLE IF NOT EXISTS "users" (
  "id" TEXT,
  "first_name" TEXT NOT NULL,
  "last_name" TEXT NOT NULL,
  "password" TEXT NOT NULL,
  "username" TEXT,
  "created_at" TIMESTAMP DEFAULT NOW(),
  PRIMARY KEY("id"),
  UNIQUE("username")
);

CREATE TABLE IF NOT EXISTS "groups" (
    "id" TEXT,
    "name" TEXT
);

-- +migrate Down
DROP TABLE IF EXISTS "users";
