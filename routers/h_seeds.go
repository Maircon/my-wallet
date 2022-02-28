package routers

import (
	"encoding/json"
	"net/http"

	"mywallet.com/db"
)

func SeedDbHandler(w http.ResponseWriter, r *http.Request) {
	db.SeedDb()
	json.NewEncoder(w).Encode("OK")
}
