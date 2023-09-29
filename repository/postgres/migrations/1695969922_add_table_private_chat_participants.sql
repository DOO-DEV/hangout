-- +migrate Up
CREATE TABLE IF NOT EXISTS "private_chat_participants" (
    "id" UUID PRIMARY KEY,
    "chat_id" UUID REFERENCES "private_chats"("id"),
    "user_id" TEXT REFERENCES "users"("id"),
    "joined_at" TIMESTAMP DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF EXISTS "private_chat_participants";