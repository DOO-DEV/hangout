-- +migrate Up

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION add_group_to_users_group()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO "users_group" ("user_id", "group_id", "joined_at", "role")
    VALUES (NEW.owner_id, NEW.id, NEW.created_at, 'admin');
    RETURN NEW;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER add_group_to_users_group
AFTER INSERT ON "groups"
FOR EACH ROW EXECUTE FUNCTION add_group_to_users_group();
-- +migrate StatementEnd


-- +migrate Down
DROP TRIGGER IF EXISTS add_group_to_users_group ON "groups";