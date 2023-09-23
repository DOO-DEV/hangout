-- +migrate Up
CREATE TABLE IF NOT EXISTS "messages" (
    "id" SERIAL PRIMARY KEY,
    "chat_id" TEXT REFERENCES "chats"("id"),
    "content" TEXT,
    "type" TEXT,
    "timestamp" TIMESTAMP DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF EXISTS "messages";