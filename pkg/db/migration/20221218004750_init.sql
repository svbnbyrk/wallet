-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "transaction" (
	id uuid DEFAULT uuid_generate_v4 (),
	wallet_id uuid,
	transaction_type INTEGER NOT NULL,
	currency VARCHAR (255) NOT NULL,
     balance DECIMAL,
     amount DECIMAL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "user" (
	id uuid DEFAULT uuid_generate_v4 (),
	name VARCHAR (255),
	email VARCHAR (255)
);

CREATE TABLE IF NOT EXISTS wallet (
	id uuid DEFAULT uuid_generate_v4 (),
	user_id uuid,
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
