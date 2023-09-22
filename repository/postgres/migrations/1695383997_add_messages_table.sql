-- +migrate Up
CREATE TABLE IF NOT EXISTS "messages" (
    "sender" TEXT,
    "receiver" TEXT,
    "content" TEXT,
    "type" TEXT,
    "timestamp" TIMESTAMP DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF EXISTS "messages";