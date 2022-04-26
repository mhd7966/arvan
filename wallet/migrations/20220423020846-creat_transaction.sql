-- +migrate Up
CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  user_id INT,
  code VARCHAR,
  code_type INT,
  value INT,
  value_type INT,
  init_balance INT,
  new_balance INT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP
);
-- +migrate Down
DROP TABLE transactions;
