package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Postgres2019!"
	dbname   = "postgres"
)

var globalSQL *sql.DB

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetDbInstance() *sql.DB {
	return globalSQL
}

func ConnectDatabase() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	if os.Getenv("DATABASE_URL") != "" {
		psqlconn = os.Getenv("DATABASE_URL")
	}

	fmt.Println(os.Getenv("DATABASE_URL") != "")

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	globalSQL = db

	// defer db.Close()

	err = db.Ping()

	CheckError(err)

	fmt.Println("Connected!")

}
