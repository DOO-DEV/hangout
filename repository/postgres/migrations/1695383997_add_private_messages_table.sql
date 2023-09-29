-- +migrate Up
CREATE TABLE IF NOT EXISTS "private_messages" (
    "id" SERIAL PRIMARY KEY,
    "chat_id" UUID REFERENCES "private_chats"("id"),
    "sender_id" TEXT REFERENCES "users"("id"),
    "content" TEXT,
    "type" INT NOT NULL,
    "status" INT NOT NULL,
    "timestamp" TIMESTAMP DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF EXISTS "messages";