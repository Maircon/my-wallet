package routers

import (
	"encoding/json"
	"net/http"

	"mywallet.com/db"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user db.User
	err := json.NewDecoder(r.Body).Decode(&user)
	CheckError(err)

	idUser := db.CreateUser(user)
	idWallet := db.CreateWalletTo(idUser)

	dtoResponse := struct {
		IdUser   string `json:"idUser"`
		IdWallet string `json:"idWallet"`
	}{
		IdUser:   idUser,
		IdWallet: idWallet,
	}

	json.NewEncoder(w).Encode(&dtoResponse)
}

func ListUsersAndWallets(w http.ResponseWriter, r *http.Request) {
	users := db.ListUsersAndWallets()
	json.NewEncoder(w).Encode(&users)
}
