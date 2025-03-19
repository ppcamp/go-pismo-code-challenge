-- 2025-03-18 19:29

CREATE SCHEMA "pismo";

CREATE TABLE IF NOT EXISTS pismo."accounts"(
    id SERIAL PRIMARY KEY,
    document_number VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS pismo."operations_types"(
    id SERIAL PRIMARY KEY,
    description VARCHAR(512) NOT NULL
);

CREATE TABLE IF NOT EXISTS pismo."transactions"(
    transaction_id SERIAL PRIMARY KEY,
    event_date TIMESTAMP DEFAULT NOW(),
    amount NUMERIC(10, 2),

    account_id  INT,
    operation_type_id INT,

    FOREIGN KEY(account_id) REFERENCES pismo.accounts(id),
    FOREIGN KEY(operation_type_id) REFERENCES pismo.operations_types(id)
);
