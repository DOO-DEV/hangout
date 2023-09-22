-- +migrate Up
CREATE TABLE IF NOT EXISTS "groups_connections" (
  "id" SERIAL PRIMARY KEY,
  "from" TEXT REFERENCES "groups"("id"),
  "to" TEXT REFERENCES  "groups"("id"),
  "created_at" TIMESTAMP DEFAULT NOW()
);

CREATE UNIQUE INDEX "dup_group_conn_1" ON "groups_connections" (LEAST("from", "to"));
CREATE UNIQUE INDEX "dup_group_conn_2" ON "groups_connections" (GREATEST("from", "to"));


-- +migrate Down
DROP TABLE IF EXISTS "groups_connections";
DROP INDEX "dup_group_conn_1" CASCADE;
DROP INDEX "dup_group_conn_2" CASCADE;