package routers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func HttpRoutersAndListener() {
	route := mux.NewRouter()
	route.HandleFunc("/categories", GetCategoriesHandler).Methods("GET")

	route.HandleFunc("/transactions", CreateTransactionHandler).Methods("POST")
	route.HandleFunc("/transactions/all", ListAllTransactionsHandler).Methods("GET")

	route.HandleFunc("/users", CreateUserHandler).Methods("POST")
	route.HandleFunc("/users/all", ListUsersAndWalletsHandler).Methods("GET")

	route.HandleFunc("/seedDb", SeedDbHandler).Methods("GET")

	port := ":8080"
	if os.Getenv("PORT") != "" {
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}
	fmt.Println(port)
	http.ListenAndServe(port, route)
}
