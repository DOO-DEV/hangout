-- +migrate Up
CREATE TABLE IF NOT EXISTS "pending_list" (
    "user_id" TEXT,
    "group_id" TEXT,
    "sent_at" TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY ("user_id", "group_id"),
    FOREIGN KEY ("user_id") REFERENCES users("id"),
    FOREIGN KEY ("group_id") REFERENCES groups("id")
);

-- +migrate Down
DROP TABLE IF EXISTS "pending_list" ;