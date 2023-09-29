-- +migrate Up
CREATE TABLE IF NOT EXISTS "group_chats" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW()
);


-- +migrate Down
DROP TABLE IF EXISTS "chats";