package routers

import (
	"encoding/json"
	"net/http"

	"mywallet.com/db"
)

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	resultado := db.GetCategories()
	json.NewEncoder(w).Encode(resultado)
}
