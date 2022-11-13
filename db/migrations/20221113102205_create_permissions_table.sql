-- migrate:up
CREATE TYPE permission_actions AS ENUM ('READ', 'WRITE');
CREATE TYPE permission_types AS ENUM ('GLOBAL', 'EXCLUSIVE');
CREATE TABLE IF NOT EXISTS permissions
(
  id VARCHAR(50) NOT NULL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  is_deleted BOOLEAN,
  name VARCHAR(50) NOT NULL,
  resource VARCHAR(50) NOT NULL,
  action permission_actions NOT NULL,
  type permission_types NOT NULL
)


-- migrate:down
DROP TABLE IF EXISTS permissions
DROP TYPE IF EXISTS permission_actions
DROP TYPE IF EXISTS permission_types

