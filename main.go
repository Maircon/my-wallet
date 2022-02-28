package main

import (
	"fmt"

	"mywallet.com/db"
	"mywallet.com/routers"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db.ConnectDatabase()
	db.Migration()
	fmt.Println("OK")
	routers.HttpRoutersAndListener()
}
