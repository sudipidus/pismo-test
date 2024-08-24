CREATE TABLE transaction (
  id SERIAL PRIMARY KEY,
  account_id INTEGER NOT NULL,
  operation_type_id INTEGER NOT NULL,
  amount DECIMAL(10, 2) NOT NULL,
  transaction_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (account_id) REFERENCES account(id),
  FOREIGN KEY (operation_type_id) REFERENCES operation_type(id)
);