-- +migrate Up
CREATE TABLE IF NOT EXISTS "group_chat_participants" (
    "id" UUID PRIMARY KEY,
    "chat_id" UUID REFERENCES "group_chats"("id"),
    "user_id" TEXT REFERENCES "users"("id"),
    "role" INT NOT NULL,
    "joined_at" TIMESTAMP DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF EXISTS "group_chat_participants";