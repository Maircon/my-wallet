package routers

import (
	"encoding/json"
	"net/http"

	"mywallet.com/db"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user db.User
	err := json.NewDecoder(r.Body).Decode(&user)
	CheckError(err)

	idUser := db.CreateUser(user)
	walletName := user.FullName + "-Wallet"
	idWallet := db.CreateWalletTo(idUser, walletName)

	dtoResponse := struct {
		IdUser   string `json:"idUser"`
		IdWallet string `json:"idWallet"`
	}{
		IdUser:   idUser,
		IdWallet: idWallet,
	}

	json.NewEncoder(w).Encode(&dtoResponse)
}

func ListUsersAndWalletsHandler(w http.ResponseWriter, r *http.Request) {
	users := db.ListUsersAndWallets()
	json.NewEncoder(w).Encode(&users)
}
