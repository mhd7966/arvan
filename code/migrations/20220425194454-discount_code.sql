-- +migrate Up
CREATE TABLE discount_codes (
  id SERIAL PRIMARY KEY,
  name TEXT,
  capacity INT,
  max_capacity INT,
  expiration_date TIMESTAMP,
  max_discount INT,
  min_purchase INT,
  percent INT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP
);
-- +migrate Down
DROP TABLE discount_codes;
