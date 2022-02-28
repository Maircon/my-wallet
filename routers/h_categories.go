package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"mywallet.com/db"
)

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	resultado := db.GetCategories()

	fmt.Println(resultado)

	json.NewEncoder(w).Encode(resultado)
}
