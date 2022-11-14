-- migrate:up
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'permission_actions') THEN
        CREATE TYPE permission_actions AS ENUM ('READ', 'WRITE');
    END IF;
     IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'permission_actions') THEN
        CREATE TYPE permission_actions AS ENUM ('READ', 'WRITE');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'permission_types') THEN
        CREATE TYPE permission_types AS ENUM ('GLOBAL', 'EXCLUSIVE');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS permissions
(
  id VARCHAR(50) NOT NULL PRIMARY KEY,
  resource VARCHAR(20) NOT NULL,
  action permission_actions NOT NULL,
  type permission_types NOT NULL
);


-- migrate:down
DROP TABLE IF EXISTS permissions;
DROP TYPE IF EXISTS permission_actions;
DROP TYPE IF EXISTS permission_types;

