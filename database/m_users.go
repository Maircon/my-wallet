package db

type User struct {
	IdUser   string `json:"userId"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}

type UserAndWallet struct {
	IdUser   string  `json:"idUser"`
	Email    string  `json:"email"`
	IdWallet string  `json:"idWallet"`
	Balance  float32 `json:"balance"`
}

func CreateUser(user User) string {
	idUser := GenerateUUID()
	_, err := GetDbInstance().Exec(`
		INSERT INTO users
			(id_user, full_name, email)
		VALUES ($1,$2,$3)
	`, idUser, user.FullName, user.Email)
	CheckError(err)

	return idUser
}

func ListUsersAndWallets() []UserAndWallet {
	rows, err := GetDbInstance().Query(`
		SELECT
			u.id_user,
			u."email",
			w.id_wallet,
			w.balance
		FROM
			users u,
			wallets w
		where
			u.id_user = w.id_user
	`)
	CheckError(err)

	var UserAndWalletList []UserAndWallet

	for rows.Next() {
		var userAndWallet UserAndWallet
		err := rows.Scan(
			&userAndWallet.IdUser,
			&userAndWallet.Email,
			&userAndWallet.IdWallet,
			&userAndWallet.Balance,
		)
		CheckError(err)

		UserAndWalletList = append(UserAndWalletList, userAndWallet)
	}

	return UserAndWalletList
}
