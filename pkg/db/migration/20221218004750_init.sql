-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS transactions (
	id SERIAL PRIMARY KEY,
	wallet_id INTEGER,
	transaction_type VARCHAR (255) NOT NULL,
	currency VARCHAR (255) NOT NULL,
     balance DECIMAL,
     amount DECIMAL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
	id  SERIAL PRIMARY KEY,
	name VARCHAR (255),
	email VARCHAR (255)
);

CREATE TABLE IF NOT EXISTS wallets (
	id  SERIAL PRIMARY KEY,
	user_id INTEGER,
	currency VARCHAR (255),
     balance DECIMAL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "transaction";
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS wallet;
-- +goose StatementEnd
