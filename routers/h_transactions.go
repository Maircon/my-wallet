package routers

import (
	"encoding/json"
	"net/http"

	"mywallet.com/db"
)

const EARN_TRANSACTION_TYPE uint8 = 2

func ListAllTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	idWallet := r.URL.Query().Get("idWallet")
	transactionsList := db.ListAllTransactionsByIdWallet(idWallet)
	json.NewEncoder(w).Encode(transactionsList)
}

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var expense db.Transaction
	var newBalance float32

	err := json.NewDecoder(r.Body).Decode(&expense)
	CheckError(err)

	idTransaction := db.CreateTransaction(expense)

	TotalBalance := db.GetTotalBalanceByIdUser(expense.IdWallet)

	if expense.IdTransactionType == EARN_TRANSACTION_TYPE {
		newBalance = TotalBalance + expense.Amount
	} else {
		newBalance = TotalBalance - expense.Amount
	}

	db.UpdateWalletBalance(
		newBalance,
		expense.IdWallet,
	)

	db.InsertAmountHistory(
		idTransaction,
		TotalBalance,
		newBalance,
	)

	dtoResponse := struct {
		Id string `json:"Id"`
	}{
		Id: idTransaction,
	}

	json.NewEncoder(w).Encode(&dtoResponse)
}
