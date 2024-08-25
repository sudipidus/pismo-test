CREATE TABLE operation_types (
                                 id SERIAL PRIMARY KEY,
                                 type VARCHAR(255) NOT NULL,
                                 description TEXT,
                                 is_credit BOOLEAN NOT NULL,
                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
