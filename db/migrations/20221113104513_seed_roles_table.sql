-- migrate:up
INSERT INTO
  roles ("name", "description")
VALUES
  ('ADMIN', 'Role which allow user full access to all resources'),
  ('USER', 'Role which restrictied access to few resources');

UPDATE users
SET role_name='USER'
WHERE role_name IS NULL;

-- migrate:down
DELETE FROM roles
WHERE name IN ('ADMIN', 'USER');

