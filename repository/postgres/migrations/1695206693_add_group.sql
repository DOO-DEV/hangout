-- +migrate Up
CREATE TABLE IF NOT EXISTS "groups" (
    "id" TEXT,
    "name" TEXT NOT NULL ,
    "owner_id" TEXT UNIQUE,
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY ("id"),
    FOREIGN KEY ("owner_id") REFERENCES users("id")
);

CREATE TABLE IF NOT EXISTS "users_group" (
    "user_id" TEXT,
    "group_id" TEXT,
    "joined_at" TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY ("user_id"),
    FOREIGN KEY ("user_id") REFERENCES users("id") ON DELETE CASCADE,
    FOREIGN KEY ("group_id") REFERENCES groups("id") ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS "groups";
DROP TABLE IF EXISTS "users_group";