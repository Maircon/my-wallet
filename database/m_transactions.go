package db

type Transaction struct {
	IdTransaction     string  `json:"idTransaction"`
	Amount            float32 `json:"amount"`
	Date              string  `json:"date"`
	IdTransactionType uint8   `json:"idTransactionType"`
	IdCategory        uint8   `json:"idCategory"`
	IdPaymentType     uint8   `json:"idPaymentType"`
	IdWallet          string  `json:"idWallet"`
	Description       string  `json:"description"`
}

type TransactionNamed struct {
	IdTransaction       string  `json:"idTransaction"`
	Amount              float32 `json:"amount"`
	Date                string  `json:"date"`
	TransactionTypeName string  `json:"TransactionTypeName"`
	CategoryName        string  `json:"CategoryName"`
	PaymentTypeName     string  `json:"PaymentTypeName"`
	IdWallet            string  `json:"idWallet"`
}

func CreateTransaction(transaction Transaction) string {
	idTransaction := GenerateUUID()
	srt := `
		INSERT INTO transactions
			(id_transaction, amount, date, id_transaction_type, id_category, id_payment_type, id_wallet, description)
		VALUES ($1,$2,$3,$4,$5,$6,$7, $8)
	`
	_, err := GetDbInstance().Exec(srt,
		idTransaction,
		transaction.Amount,
		transaction.Date,
		transaction.IdTransactionType,
		transaction.IdCategory,
		transaction.IdPaymentType,
		transaction.IdWallet,
		transaction.Description,
	)

	CheckError(err)
	return idTransaction
}

func ListAllTransactionsByIdWallet(idWallet string) []TransactionNamed {
	transactionList := []TransactionNamed{}
	rows, err := GetDbInstance().Query(`
		SELECT
			t.id_transaction,
			t.amount,
			t."date",
			tt."name" "transaction_type",
			c."name" "category",
			pt."name" "payment_type",
			t.id_wallet
		FROM transactions t
		JOIN
			transaction_types tt ON tt.id_transaction_type = t.id_transaction_type
		JOIN
			categories c ON c.id_category = t.id_category
		JOIN
			payment_types pt ON pt.id_payment_type = t.id_payment_type
		WHERE t.id_wallet = $1
	`, idWallet)

	CheckError(err)

	for rows.Next() {
		transaction := TransactionNamed{}
		err := rows.Scan(
			&transaction.IdTransaction,
			&transaction.Amount,
			&transaction.Date,
			&transaction.TransactionTypeName,
			&transaction.CategoryName,
			&transaction.PaymentTypeName,
			&transaction.IdWallet,
		)

		CheckError(err)

		transactionList = append(transactionList, transaction)
	}

	return transactionList
}
