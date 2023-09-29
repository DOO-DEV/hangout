-- +migrate Up
CREATE TABLE IF NOT EXISTS "messages" (
    "id" SERIAL PRIMARY KEY,
    "chat_id" UUID REFERENCES "chats"("id"),
    "sender_id" TEXT REFERENCES "users"("id"),
    "content" TEXT,
    "type" INT NOT NULL,
    "timestamp" TIMESTAMP DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF EXISTS "messages";