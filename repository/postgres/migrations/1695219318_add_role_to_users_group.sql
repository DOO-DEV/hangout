-- +migrate Up
ALTER TABLE "users_group" ADD COLUMN "role" TEXT CHECK (role IN ('admin', 'normal'));

-- +migrate Down
ALTER TABLE "users_group" DROP COLUMN "role";