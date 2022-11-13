-- migrate:up
CREATE TABLE IF NOT EXISTS roles
(
  name VARCHAR(50) NOT NULL PRIMARY KEY,
  description TEXT
)


-- migrate:down
DROP TABLE IF EXISTS roles

