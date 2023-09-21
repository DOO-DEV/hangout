-- +migrate Up
CREATE TABLE IF NOT EXISTS "groups_connections" (
  "id" SERIAL PRIMARY KEY,
  "from" TEXT REFERENCES "groups"("id"),
  "to" TEXT REFERENCES  "groups"("id"),
  "created_at" TIMESTAMP DEFAULT NOW()
);

CREATE UNIQUE INDEX unique_group_connection ON
"groups_connections" (LEAST("id"::TEXT, "from"::TEXT), GREATEST("id"::TEXT, "from"::TEXT));


-- +migrate Down
DROP TABLE IF EXISTS "groups_connections";
DROP INDEX unique_group_connection;