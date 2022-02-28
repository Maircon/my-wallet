package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"mywallet.com/db"
)

func ListAllExpenses(w http.ResponseWriter, r *http.Request) {
	expenses := []struct {
		Title string `json:"title"`
		Value string `json:"value"`
	}{
		{
			Title: "Maircon",
			Value: "dotinha",
		},
		{
			Title: "Maircon",
			Value: "dotinha",
		},
	}
	fmt.Println(expenses)
	// fmt.Fprintf(w, "Welcome Home, %q", html.EscapeString(r.URL.Path))
	json.NewEncoder(w).Encode(expenses)
}

func InsertAnExpense(w http.ResponseWriter, r *http.Request) {
	var expense db.Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	CheckError(err)

	idTransaction := db.InsertExpense(expense)

	TotalBalance := db.GetTotalBalanceByIdUser(expense.IdWallet)
	newBalance := TotalBalance - expense.Amount

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
