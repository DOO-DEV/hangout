-- +migrate Up
CREATE TABLE IF NOT EXISTS "messages" (
    "id" SERIAL PRIMARY KEY,
    "sender" TEXT REFERENCES "users"("id"),
    "receiver" TEXT REFERENCES "users"("id"),
    "content" TEXT,
    "type" TEXT,
    "timestamp" TIMESTAMP DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF EXISTS "messages";