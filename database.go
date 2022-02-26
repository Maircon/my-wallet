package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Postgres2019!"
	dbname   = "postgres"
)

var globalSQL *sql.DB

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GenerateUUID() string {
	return uuid.NewString()
}

func ConnectDatabase() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	if os.Getenv("DATABASE_URL") != "" {
		psqlconn = os.Getenv("DATABASE_URL")
	}

	fmt.Println(os.Getenv("DATABASE_URL") != "")

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	globalSQL = db

	// defer db.Close()

	err = db.Ping()

	CheckError(err)

	fmt.Println("Connected!")
}

func GetCategories() []Categories {
	rows, err := globalSQL.Query("SELECT * FROM categories")
	CheckError(err)
	var res []Categories
	for rows.Next() {
		var category Categories
		err = rows.Scan(&category.IdCategory, &category.Name)
		CheckError(err)
		if category.IdCategory == 4 {
			teste := res[1]
			teste.IdCategory = 55
		}
		res = append(res, category)
	}
	// fmt.Println(res)
	return res
}

func InsertExpense(expense Expense) string {
	idTransaction := GenerateUUID()
	srt := `
		INSERT INTO transactions
			(id_transaction, amount, date, id_transaction_type, id_category, id_payment_type, id_wallet)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
	`
	_, err := globalSQL.Query(srt,
		idTransaction,
		expense.Amount,
		expense.Date,
		expense.IdTransactionType,
		expense.IdCategory,
		expense.IdPaymentType,
		expense.IdWallet,
	)

	CheckError(err)
	return idTransaction
}

func GetTotalBalanceByIdUser(idWallet string) float32 {
	var balance float32
	err := globalSQL.QueryRow(`
		SELECT balance FROM
			wallets
		WHERE
			id_wallet = $1`, idWallet).Scan(&balance)
	CheckError(err)
	return balance
}

func UpdateWalletBalance(balance float32, idWallet string) {
	_, err := globalSQL.Query(`
		UPDATE wallets
		SET balance = $1
		WHERE id_wallet = $2
	`, balance, idWallet)

	CheckError(err)
}

func InsertAmountHistory(idTransaction string, lastAmount float32, nextAmount float32) {
	srt := `
		INSERT INTO amount_history
			(id_transaction, last_amount, next_amount)
		VALUES ($1,$2,$3)
	`
	_, err := globalSQL.Query(srt,
		idTransaction,
		lastAmount,
		nextAmount,
	)

	CheckError(err)
}

func migration() {
	_, err := globalSQL.Exec(`
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

	_, err = globalSQL.Exec(`
		CREATE TABLE IF NOT EXISTS transaction_types (
			id_transaction_type SMALLSERIAL not null,
			"name" VARCHAR NOT NULL,
			PRIMARY KEY (id_transaction_type)
		)
	`)

	CheckError(err)
	fmt.Println("transaction_types created or exists")

	_, err = globalSQL.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id_category SMALLSERIAL NOT NULL,
			"name" VARCHAR NOT NULL,
			PRIMARY KEY (id_category)
		)
	`)

	CheckError(err)
	fmt.Println("categories created or exists")

	_, err = globalSQL.Exec(`
		CREATE TABLE IF NOT EXISTS payment_types (
			id_payment_type SMALLSERIAL NOT NULL,
			"name" VARCHAR NOT NULL,
			PRIMARY KEY (id_payment_type)
		)
	`)

	CheckError(err)
	fmt.Println("payment_types created or exists")

	_, err = globalSQL.Exec(`
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

	_, err = globalSQL.Exec(`
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

	_, err = globalSQL.Exec(`
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

func seedDb() {
	_, err := globalSQL.Exec(`
		INSERT INTO public.categories ("name")
			VALUES ('Housing');
		INSERT INTO public.categories ("name")
			VALUES ('Transportation');
		INSERT INTO public.categories ("name")
			VALUES ('Supermarket');
		INSERT INTO public.categories ("name")
			VALUES ('Medical & Healthcare');
		INSERT INTO public.categories ("name")
			VALUES ('Clothes');
		INSERT INTO public.categories ("name")
			VALUES ('Bars & Restaurants');
		INSERT INTO public.categories ("name")
			VALUES ('e-Food');
		INSERT INTO public.categories ("name")
			VALUES ('Recreation & Entertainment');
		INSERT INTO public.categories ("name")
			VALUES ('Travel');
		INSERT INTO public.categories ("name")
			VALUES ('Subscriptions & Services');
		INSERT INTO public.categories ("name")
			VALUES ('Work');
		INSERT INTO public.categories ("name")
			VALUES ('Salary');
		INSERT INTO public.categories ("name")
			VALUES ('Family');
		INSERT INTO public.categories ("name")
			VALUES ('Personal Care');
	`)

	CheckError(err)
	fmt.Println("categories seed ok")

	_, err = globalSQL.Exec(`
		INSERT INTO public.payment_types ("name")
			VALUES ('cash');
		INSERT INTO public.payment_types ("name")
			VALUES ('credit');
		INSERT INTO public.payment_types ("name")
			VALUES ('loan');
		INSERT INTO public.payment_types ("name")
			VALUES ('meal_ticket');
	`)

	CheckError(err)
	fmt.Println("payment_types seed ok")

	_, err = globalSQL.Exec(`
		INSERT INTO public.transaction_types ("name")
			VALUES ('expense');
		INSERT INTO public.transaction_types ("name")
			VALUES ('earn');
	`)

	CheckError(err)
	fmt.Println("transaction_types seed ok")
}
