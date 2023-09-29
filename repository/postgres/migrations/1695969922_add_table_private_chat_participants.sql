-- +migrate UP
CREATE TABLE IF NOT EXISTS "private_chat_participants" (
    "id" UUID,
    "chat_id" UUID REFERENCES "chats"("id"),
    "user_id" TEXT REFERENCES "users"("id"),
    "joined_at" TIMESTAMP DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF EXISTS "private_chat_participants";