package db

import "fmt"

func SeedDb() {
	_, err := GetDbInstance().Exec(`
		INSERT INTO public.categories ("name")
			VALUES ('Housing');
		INSERT INTO public.categories ("name")
			VALUES ('Transportation');
		INSERT INTO public.categories ("name")
			VALUES ('Supermarket');
		INSERT INTO public.categories ("name")
			VALUES ('Medical & Healthcare');
		INSERT INTO public.categories ("name")
			VALUES ('Shopping');
		INSERT INTO public.categories ("name")
			VALUES ('Bars & Restaurants');
		INSERT INTO public.categories ("name")
			VALUES ('e-Food');
		INSERT INTO public.categories ("name")
			VALUES ('Recreation & Entertainment');
		INSERT INTO public.categories ("name")
			VALUES ('Travel');
		INSERT INTO public.categories ("name")
			VALUES ('Subscriptions & Services');
		INSERT INTO public.categories ("name")
			VALUES ('Work');
		INSERT INTO public.categories ("name")
			VALUES ('Salary');
		INSERT INTO public.categories ("name")
			VALUES ('Family');
		INSERT INTO public.categories ("name")
			VALUES ('Personal Care');
		INSERT INTO public.categories ("name")
			VALUES ('Animal Meal');
	`)

	CheckError(err)
	fmt.Println("categories seed ok")

	_, err = GetDbInstance().Exec(`
		INSERT INTO public.payment_types ("name")
			VALUES ('cash');
		INSERT INTO public.payment_types ("name")
			VALUES ('credit');
		INSERT INTO public.payment_types ("name")
			VALUES ('loan');
		INSERT INTO public.payment_types ("name")
			VALUES ('meal_ticket');
	`)

	CheckError(err)
	fmt.Println("payment_types seed ok")

	_, err = GetDbInstance().Exec(`
		INSERT INTO public.transaction_types ("name")
			VALUES ('expense');
		INSERT INTO public.transaction_types ("name")
			VALUES ('earn');
	`)

	CheckError(err)
	fmt.Println("transaction_types seed ok")
}
