package db

type Transaction struct {
	IdTransaction     string  `json:"idTransaction"`
	Amount            float32 `json:"amount"`
	Date              string  `json:"date"`
	IdTransactionType uint8   `json:"idTransactionType"`
	IdCategory        uint8   `json:"idCategory"`
	IdPaymentType     uint8   `json:"idPaymentType"`
	IdWallet          string  `json:"idWallet"`
}

func CreateTransaction(transaction Transaction) string {
	idTransaction := GenerateUUID()
	srt := `
		INSERT INTO transactions
			(id_transaction, amount, date, id_transaction_type, id_category, id_payment_type, id_wallet)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
	`
	_, err := GetDbInstance().Query(srt,
		idTransaction,
		transaction.Amount,
		transaction.Date,
		transaction.IdTransactionType,
		transaction.IdCategory,
		transaction.IdPaymentType,
		transaction.IdWallet,
	)

	CheckError(err)
	return idTransaction
}

func ListAllTransactionsByIdWallet(idWallet string) []Transaction {
	transactionList := []Transaction{}
	return transactionList
}
