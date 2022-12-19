-- +goose Up
-- +goose StatementBegin

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

CREATE TABLE IF NOT EXISTS exchanges (
	id  SERIAL PRIMARY KEY,
	currency VARCHAR (255),
	rate VARCHAR (255)
);

INSERT INTO exchanges (currency, rate) VALUES('USD', 1);
INSERT INTO exchanges (currency, rate) VALUES('GBP', 0.82089785506879);
INSERT INTO exchanges (currency, rate) VALUES('TRY', 18.64843462861);
INSERT INTO exchanges (currency, rate) VALUES('RUB', 66.907120956442);
INSERT INTO exchanges (currency, rate) VALUES('GBP', 0.82089785506879);
INSERT INTO exchanges (currency, rate) VALUES('CHF', 0.92960718755693);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS wallets;
DROP TABLE IF EXISTS exchanges;
-- +goose StatementEnd
