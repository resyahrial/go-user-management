-- migrate:up
CREATE TABLE IF NOT EXISTS roles_permissions
(
  id VARCHAR(100) NOT NULL PRIMARY KEY,
  role_name VARCHAR(50) NOT NULL,
  permission_id VARCHAR(50) NOT NULL
);

ALTER TABLE roles_permissions
ADD CONSTRAINT fk_name_role_name
FOREIGN KEY (role_name)
REFERENCES roles (name)
ON DELETE CASCADE;

ALTER TABLE roles_permissions
ADD CONSTRAINT fk_id_permission_id
FOREIGN KEY (permission_id)
REFERENCES permissions (id)
ON DELETE CASCADE;

-- migrate:down
ALTER TABLE roles_permissions
DROP CONSTRAINT fk_id_permission_id;

ALTER TABLE roles_permissions
DROP CONSTRAINT fk_name_role_name;

DROP TABLE IF EXISTS role_permissions;

