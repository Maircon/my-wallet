package db

type Categories struct {
	IdCategory int    `json:"idCategory"`
	Name       string `json:"name"`
}

func GetCategories() []Categories {
	rows, err := GetDbInstance().Query("SELECT * FROM categories")
	CheckError(err)
	var res []Categories
	for rows.Next() {
		var category Categories
		err = rows.Scan(&category.IdCategory, &category.Name)
		CheckError(err)
		// if category.IdCategory == 4 {
		// 	teste := res
		// 	teste[1].IdCategory = 55
		// }
		res = append(res, category)
	}
	return res
}
