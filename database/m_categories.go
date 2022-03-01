package db

type Categories struct {
	IdCategory int    `json:"idCategory"`
	Name       string `json:"name"`
}

func GetCategories() []Categories {
	rows, err := GetDbInstance().Query("SELECT * FROM categories")
	CheckError(err)
	var categoriesList []Categories

	for rows.Next() {
		var category Categories
		err = rows.Scan(
			&category.IdCategory,
			&category.Name,
		)

		CheckError(err)

		categoriesList = append(categoriesList, category)
	}
	return categoriesList
}
