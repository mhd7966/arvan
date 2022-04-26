-- +migrate Up
CREATE TABLE charge_codes (
  id SERIAL PRIMARY KEY,
  name TEXT,
  value INT,
  capacity INT,
  max_capacity INT,
  expiration_date TIMESTAMP,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP
);
-- +migrate Down
DROP TABLE charge_codes;