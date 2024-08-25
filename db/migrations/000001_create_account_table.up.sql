CREATE TABLE accounts (
                          id  SERIAL PRIMARY KEY,
                          document_number VARCHAR(255) NOT NULL,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
