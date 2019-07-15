package main

import (
	. "GoProject/JsonFetching"
	. "GoProject/dbReadCreate"
	. "GoProject/routes"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Starting Go mysql connection")
	db := DbConn()
	fmt.Println("connection made ", db)

	FetchFromJson()

	CreateRoutes()
	defer db.Close()
}
