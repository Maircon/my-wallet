package db

const DEFAULT_BALANCE = 0

func GetTotalBalanceByIdUser(idWallet string) float32 {
	var balance float32
	err := GetDbInstance().QueryRow(`
		SELECT balance FROM
			wallets
		WHERE
			id_wallet = $1`, idWallet).Scan(&balance)
	CheckError(err)
	return balance
}

func UpdateWalletBalance(balance float32, idWallet string) {
	_, err := GetDbInstance().Exec(`
		UPDATE wallets
		SET balance = $1
		WHERE id_wallet = $2
	`, balance, idWallet)

	CheckError(err)
}

func CreateWalletTo(idUser string, walletName string) string {
	idWallet := GenerateUUID()
	_, err := GetDbInstance().Exec(`
		INSERT INTO wallets
			(id_wallet, id_user, name, balance)
		VALUES ($1,$2,$3, $4)
	`, idWallet, idUser, walletName, DEFAULT_BALANCE)

	CheckError(err)

	return idWallet
}
