package db

import "fmt"

func Migration() {
	_, err := GetDbInstance().Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id_user UUID NOT NULL,
			full_name VARCHAR NOT NULL,
			email VARCHAR NOT NULL,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP,
			deleted_at TIMESTAMP,
			PRIMARY KEY (id_user)
		)
	`)

	CheckError(err)
	fmt.Println("users created or exists")

	_, err = GetDbInstance().Exec(`
		CREATE TABLE IF NOT EXISTS transaction_types (
			id_transaction_type SMALLSERIAL not null,
			"name" VARCHAR NOT NULL,
			PRIMARY KEY (id_transaction_type)
		)
	`)

	CheckError(err)
	fmt.Println("transaction_types created or exists")

	_, err = GetDbInstance().Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id_category SMALLSERIAL NOT NULL,
			"name" VARCHAR NOT NULL,
			PRIMARY KEY (id_category)
		)
	`)

	CheckError(err)
	fmt.Println("categories created or exists")

	_, err = GetDbInstance().Exec(`
		CREATE TABLE IF NOT EXISTS payment_types (
			id_payment_type SMALLSERIAL NOT NULL,
			"name" VARCHAR NOT NULL,
			PRIMARY KEY (id_payment_type)
		)
	`)

	CheckError(err)
	fmt.Println("payment_types created or exists")

	_, err = GetDbInstance().Exec(`
		CREATE TABLE IF NOT EXISTS wallets (
			id_wallet UUID NOT NULL,
			id_user UUID NOT NULL,
			"name" VARCHAR NOT NULL,
			balance NUMERIC(12,2),
			PRIMARY KEY (id_wallet),
			CONSTRAINT fk_id_user FOREIGN KEY(id_user) REFERENCES users(id_user)
		)
	`)

	CheckError(err)
	fmt.Println("wallets created or exists")

	_, err = GetDbInstance().Exec(`
		CREATE TABLE IF NOT EXISTS transactions (
			id_transaction UUID NOT NULL,
			amount NUMERIC(12,2),
			"date" TIMESTAMP,
			id_transaction_type SMALLSERIAL,
			id_category SMALLSERIAL,
			id_payment_type SMALLSERIAL,
			id_wallet UUID,
			PRIMARY KEY (id_transaction),
			CONSTRAINT fk_id_transaction_type FOREIGN KEY(id_transaction_type) REFERENCES transaction_types(id_transaction_type),
			CONSTRAINT fk_id_category FOREIGN KEY(id_category) REFERENCES categories(id_category),
			CONSTRAINT fk_id_payment_type FOREIGN KEY(id_payment_type) REFERENCES payment_types(id_payment_type),
			CONSTRAINT fk_id_wallet FOREIGN KEY(id_wallet) REFERENCES wallets(id_wallet)
		)
	`)

	CheckError(err)
	fmt.Println("transactions created or exists")

	_, err = GetDbInstance().Exec(`
		CREATE TABLE IF NOT EXISTS amount_history (
			id_amount_history SMALLSERIAL NOT NULL,
			id_transaction UUID NOT NULL,
			last_amount NUMERIC(12,2),
			next_amount NUMERIC(12,2),
			CONSTRAINT fk_id_transaction FOREIGN KEY(id_transaction) REFERENCES transactions(id_transaction)
		)
	`)

	CheckError(err)
	fmt.Println("amount_history created or exists")
}
