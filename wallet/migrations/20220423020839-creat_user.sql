-- +migrate Up
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  phone_number TEXT,
  balance INT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP
);
-- +migrate Down
DROP TABLE users;