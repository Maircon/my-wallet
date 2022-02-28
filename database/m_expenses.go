package db

type Expense struct {
	IdTransaction     string  `json:"idTransaction"`
	Amount            float32 `json:"amount"`
	Date              string  `json:"date"`
	IdTransactionType uint8   `json:"idTransactionType"`
	IdCategory        uint8   `json:"idCategory"`
	IdPaymentType     uint8   `json:"idPaymentType"`
	IdWallet          string  `json:"idWallet"`
}

func InsertExpense(expense Expense) string {
	idTransaction := GenerateUUID()
	srt := `
		INSERT INTO transactions
			(id_transaction, amount, date, id_transaction_type, id_category, id_payment_type, id_wallet)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
	`
	_, err := GetDbInstance().Query(srt,
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

func ListAllExpensesByIdWallet(idWallet string) {}
