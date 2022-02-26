package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Categories struct {
	IdCategory int    `json:"idCategory"`
	Name       string `json:"name"`
}

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	resultado := GetCategories()

	fmt.Println(resultado)

	json.NewEncoder(w).Encode(resultado)
}

func seedDbHandler(w http.ResponseWriter, r *http.Request) {
	seedDb()
	json.NewEncoder(w).Encode("OK")
}

func ListAllExpenses(w http.ResponseWriter, r *http.Request) {
	expenses := []Dota{
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

type Expense struct {
	IdTransaction     string  `json:"idTransaction"`
	Amount            float32 `json:"amount"`
	Date              string  `json:"date"`
	IdTransactionType uint8   `json:"idTransactionType"`
	IdCategory        uint8   `json:"idCategory"`
	IdPaymentType     uint8   `json:"idPaymentType"`
	IdWallet          string  `json:"idWallet"`
}

type dtoResponse struct {
	Id string `json:"dshaudhsau"`
}

func InsertAnExpense(w http.ResponseWriter, r *http.Request) {
	var expense Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	CheckError(err)

	idTransaction := InsertExpense(expense)

	TotalBalance := GetTotalBalanceByIdUser(expense.IdWallet)
	newBalance := TotalBalance - expense.Amount

	UpdateBalanceByIdWallet(
		newBalance,
		expense.IdWallet,
	)

	CreateAmountHistory(
		idTransaction,
		TotalBalance,
		newBalance,
	)

	dtoResponse := dtoResponse{
		Id: idTransaction,
	}

	json.NewEncoder(w).Encode(&dtoResponse)
}

func UpdateBalanceByIdWallet(balance float32, idWallet string) {
	UpdateWalletBalance(balance, idWallet)
}

func CreateAmountHistory(idTransaction string, TotalBalance float32, newBalance float32) {
	InsertAmountHistory(
		idTransaction,
		TotalBalance,
		newBalance,
	)
}

func GetWalletAmount(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(1000)
}
