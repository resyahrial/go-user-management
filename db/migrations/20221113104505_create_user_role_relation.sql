-- migrate:up
ALTER TABLE users
ADD COLUMN role_name VARCHAR(50);

ALTER TABLE users
ADD CONSTRAINT fk_name_role_name
FOREIGN KEY (role_name)
REFERENCES roles (name)
ON DELETE SET NULL;

-- migrate:down
ALTER TABLE users
DROP CONSTRAINT fk_name_role_name;

ALTER TABLE users
DROP COLUMN role_name;

