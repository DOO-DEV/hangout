-- +migrate Up
CREATE TABLE IF NOT EXISTS "account_images" (
    "id" UUID PRIMARY KEY,
    "account" TEXT REFERENCES "users"("id"),
    "url" VARCHAR(256),
    "is_primary" BOOLEAN DEFAULT FALSE,
    "order" INT,
    "created_at" TIMESTAMP DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF EXISTS "account_images";