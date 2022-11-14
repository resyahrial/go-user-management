-- migrate:up
INSERT INTO
  permissions ("id", "resource", "action", "type")
VALUES
  ('users_read_global', 'users', 'READ', 'GLOBAL'),
  ('users_read_exclusive', 'users', 'READ', 'EXCLUSIVE'),
  ('users_write_global', 'users', 'WRITE', 'GLOBAL'),
  ('users_write_exclusive', 'users', 'WRITE', 'EXCLUSIVE')
ON CONFLICT DO NOTHING;

INSERT INTO 
  roles_permissions ("id", "role_name", "permission_id")
VALUES
  ('admin_users_read_global', 'ADMIN', 'users_read_global'),
  ('admin_users_write_global', 'ADMIN', 'users_write_global'),
  ('admin_users_read_exclusive', 'USER', 'users_read_exclusive'),
  ('admin_users_write_exclusive', 'USER', 'users_write_exclusive')
ON CONFLICT DO NOTHING;

-- migrate:down
DELETE FROM permissions
WHERE resource IN ('users');
