CREATE TABLE account (
  id SERIAL PRIMARY KEY,
  account_number VARCHAR(20) NOT NULL,
  account_holder VARCHAR(100) NOT NULL,
  balance DECIMAL(10, 2) NOT NULL DEFAULT 0.00
);