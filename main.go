package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Dota struct {
	Title string `json:"Title"`
	Value string `json:"Value"`
}

func httpRoutersAndListener() {
	route := mux.NewRouter()
	route.HandleFunc("/", ListAllExpenses).Methods("GET")
	route.HandleFunc("/categories", GetCategoriesHandler).Methods("GET")
	route.HandleFunc("/expense", InsertAnExpense).Methods("POST")
	route.HandleFunc("/seedDb", seedDbHandler).Methods("GET")

	port := ":8080"
	if os.Getenv("PORT") != "" {
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}
	fmt.Println(port)
	http.ListenAndServe(port, route)
}

func main() {
	ConnectDatabase()
	migration()
	fmt.Println("OK")
	httpRoutersAndListener()
}
