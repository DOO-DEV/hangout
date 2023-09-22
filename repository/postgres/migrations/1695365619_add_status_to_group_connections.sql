-- +migrate Up
ALTER TABLE "groups_connections" ADD COLUMN "accept" BOOLEAN DEFAULT FALSE;

-- +migrate Down
ALTER TABLE "groups_connections" DROP COLUMN "accept";