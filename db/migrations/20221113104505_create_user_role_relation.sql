-- migrate:up
ALTER TABLE users
ADD COLUMN IF NOT EXISTS role_name VARCHAR(50);

ALTER TABLE users
DROP CONSTRAINT IF EXISTS fk_name_role_name;

ALTER TABLE users
ADD CONSTRAINT fk_name_role_name
FOREIGN KEY (role_name)
REFERENCES roles (name)
ON DELETE SET NULL;

-- migrate:down
ALTER TABLE users
DROP CONSTRAINT IF EXISTS fk_name_role_name;

ALTER TABLE users
DROP COLUMN IF EXISTS role_name;

