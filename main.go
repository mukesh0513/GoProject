package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	Id      int
	Name    string
	Email   string
	Phone   string
	Detail1 string
	Detail2 string
}

var dbDriver = "mysql"
var dbUser = "root"
var dbPass = "root"
var dbName = "firstgodb"
var tbName = "logGenerate"

func dbConn() (db *sql.DB) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func uploadData(db *sql.DB) {

	insert, err := db.Query("INSERT INTO" + " " + tbName + "(name, email, phone, detail1, detail2) " +
		"VALUES( 'mukesh', 'mueksh@gmnail.com', '12479273', 'no detail', 'no detail again' )")

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func readData(db *sql.DB) {

	retrieve, err := db.Query("SELECT * FROM " + tbName)

	if err != nil {
		panic(err.Error())
	}

	for retrieve.Next() {
		var uid int
		var name string
		var email string
		var phone string
		var detail1 string
		var detail2 string
		err = retrieve.Scan(&uid, &name, &email, &phone, &detail1, &detail2)

		checkErr(err)
		fmt.Println(name)
		fmt.Println(email)
		fmt.Println(phone)
		fmt.Println(detail1)
		fmt.Println(detail2)

	}

	fmt.Println("Retrieving data", retrieve)
	defer retrieve.Close()
}

func main() {
	fmt.Println("Starting Go mysql connection")
	db := dbConn()
	fmt.Println("connection made ", db)

	uploadData(db)
	readData(db)
	defer db.Close()
}

func checkErr(Error error) {
	if Error != nil {
		panic(Error.Error())
	}
}
