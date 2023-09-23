-- +migrate Up
CREATE TABLE IF NOT EXISTS "chats" (
    "id" TEXT,
    "user_1" TEXT,
    "user_2" TEXT,
    "type" TEXT,
    "created_at" TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY ("id"),
    FOREIGN KEY ("user_1") REFERENCES "users"("id"),
    FOREIGN KEY ("user_2") REFERENCES "users"("id")
);


-- +migrate Down
DROP TABLE IF EXISTS "chats";