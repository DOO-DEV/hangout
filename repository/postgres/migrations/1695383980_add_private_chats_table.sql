-- +migrate Up
CREATE TABLE IF NOT EXISTS "private_chats" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL UNIQUE, --"userid-userid"
    "created_at" TIMESTAMP DEFAULT NOW()
);


-- +migrate Down
DROP TABLE IF EXISTS "chats";