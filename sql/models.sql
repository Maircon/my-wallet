CREATE TABLE users (
  id_user UUID NOT NULL,
  full_name VARCHAR NOT NULL,
  email VARCHAR NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id_user)
)

CREATE TABLE transaction_types (
  id_transaction_type SMALLSERIAL not null,
  "name" VARCHAR NOT NULL,
  PRIMARY KEY (id_transaction_type)
)

CREATE TABLE categories (
  id_category SMALLSERIAL NOT NULL,
  "name" VARCHAR NOT NULL,
  PRIMARY KEY (id_category)
)

CREATE TABLE payment_types (
  id_payment_type SMALLSERIAL NOT NULL,
  "name" VARCHAR NOT NULL,
  PRIMARY KEY (id_payment_type)
)

CREATE TABLE wallets (
  id_wallet UUID NOT NULL,
  id_user UUID NOT NULL,
  "name" VARCHAR NOT NULL,
  balance NUMERIC(12,2),
  PRIMARY KEY (id_wallet),
  CONSTRAINT fk_id_user FOREIGN KEY(id_user) REFERENCES users(id_user)
)

CREATE TABLE transactions (
  id_transaction UUID NOT NULL,
  amount NUMERIC(12,2),
  "date" TIMESTAMP,
  id_transaction_type SMALLSERIAL,
  id_category SMALLSERIAL,
  id_payment_type SMALLSERIAL,
  id_wallet UUID,
  PRIMARY KEY (id_transaction),
  CONSTRAINT fk_id_transaction_type FOREIGN KEY(id_transaction_type) REFERENCES transaction_types(id_transaction_type)
  CONSTRAINT fk_id_category FOREIGN KEY(id_category) REFERENCES categories(id_category),
  CONSTRAINT fk_id_payment_type FOREIGN KEY(id_payment_type) REFERENCES payment_types(id_payment_type),
  CONSTRAINT fk_id_wallet FOREIGN KEY(id_wallet) REFERENCES wallets(id_wallet)
)

CREATE TABLE amount_history (
  id_amount_history SERIAL PRIMARY KEY,
  id_transaction UUID NOT NULL,
  last_amount NUMERIC(12,2),
  next_amount NUMERIC(12,2),
  CONSTRAINT fk_id_transaction FOREIGN KEY(id_transaction) REFERENCES transactions(id_transaction)
)