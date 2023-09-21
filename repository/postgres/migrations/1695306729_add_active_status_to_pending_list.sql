-- +migrate UP
ALTER TABLE "pending_list" ADD COLUMN "active" BOOLEAN DEFAULT TRUE;

-- +migrate Down
ALTER TABLE "pending_list" DROP COLUMN "active";