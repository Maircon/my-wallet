package db

func InsertAmountHistory(idTransaction string, lastAmount float32, nextAmount float32) {
	srt := `
		INSERT INTO amount_history
			(id_transaction, last_amount, next_amount)
		VALUES ($1,$2,$3)
	`
	_, err := GetDbInstance().Query(srt,
		idTransaction,
		lastAmount,
		nextAmount,
	)

	CheckError(err)
}
