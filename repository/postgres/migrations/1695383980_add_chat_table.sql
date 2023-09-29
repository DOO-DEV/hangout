-- +migrate Up
CREATE TABLE IF NOT EXISTS "chats" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(80),
    "type" INT NOT NULL,
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW()
);


-- +migrate Down
DROP TABLE IF EXISTS "chats";