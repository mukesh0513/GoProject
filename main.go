package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Users struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Detail1 string `json:"detail1"`
	Detail2 string `json:"detail2"`
}

var dbDriver = "mysql"
var dbUser = "root"
var dbPass = "root"
var dbName = "firstgodb"
var tbName = "logGenerate"

func main() {
	fmt.Println("Starting Go mysql connection")
	db := dbConn()
	fmt.Println("connection made ", db)

	fetchFromJson()

	createRoutes()
	defer db.Close()
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func uploadData(user []Users) {

	db := dbConn()
	for i := 0; i < len(user); i++ {

		insert, err := db.Query("INSERT INTO" + " " + tbName + "(name, email, phone, detail1, detail2) " +
			"VALUES( '" + user[i].Name + "', '" + user[i].Email + "', '" + user[i].Phone + "', '" + user[i].Detail1 + "', '" +
			user[i].Detail2 + "' )")

		checkErr(err)
		defer insert.Close()
	}
}

func fetchFromJson() {

	var user []Users
	jsonFile, err := os.Open("log.json")
	checkErr(err)

	fmt.Println("Successfully Opened users.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &user)
	uploadData(user)
	defer jsonFile.Close()
}

func createRoutes() {

	// http.HandleFunc("/user", showUser)
	// http.HandleFunc("/admin", showAdmin)
	// http.HandleFunc("/read/{id}/", readData)
	// err := http.ListenAndServe(":9090", nil)
	// checkErr(err)
	r := mux.NewRouter()
	r.HandleFunc("/{person}/{name}", personData)
	r.HandleFunc("/{person}", fullData)
	r.HandleFunc("/{person}/{type}/{value}", showData)

	err := http.ListenAndServe(":8080", r)
	checkErr(err)
}

func personData(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	vars := mux.Vars(r)

	retrieve, err := db.Query("SELECT * FROM " + tbName + " WHERE name = '" + vars["name"] + "'")
	checkErr(err)

	if vars["person"] == "user" {
		showUser(w, retrieve)
	} else if vars["person"] == "admin" {
		showAdmin(w, retrieve)
	} else {
		panic(err.Error())
	}
	defer retrieve.Close()
}

func fullData(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	vars := mux.Vars(r)

	retrieve, err := db.Query("SELECT * FROM " + tbName)
	checkErr(err)

	if vars["person"] == "user" {
		showUser(w, retrieve)
	} else if vars["person"] == "admin" {
		showAdmin(w, retrieve)
	} else {
		panic(err.Error())
	}
	defer retrieve.Close()
}

func showUser(w http.ResponseWriter, retrieve *sql.Rows) {

	if retrieve.Next() {
		fmt.Fprintf(w, "Retrieving data\n")
		for retrieve.Next() {
			var uid int
			var name string
			var email string
			var phone string
			var detail1 string
			var detail2 string
			err := retrieve.Scan(&uid, &name, &email, &phone, &detail1, &detail2)

			checkErr(err)
			fmt.Fprintf(w, "name:\t"+name)
			fmt.Fprintf(w, "\t\temail:\t"+"xxxxxxxxxxxxxx")
			fmt.Fprintf(w, "\t\tphone:\t"+"xxxxxxxxxxxxxx")
			fmt.Fprintf(w, "\t\tdetail1:\t"+detail1)
			fmt.Fprintf(w, "\t\tdetail2:\t"+detail2)
			fmt.Fprintf(w, "\n")
		}
	} else {
		fmt.Fprintf(w, "User not exist!")
	}

}

func showAdmin(w http.ResponseWriter, retrieve *sql.Rows) {

	if retrieve.Next() {
		fmt.Fprintf(w, "Retrieving data\n")
		for retrieve.Next() {
			var uid int
			var name string
			var email string
			var phone string
			var detail1 string
			var detail2 string
			err := retrieve.Scan(&uid, &name, &email, &phone, &detail1, &detail2)

			checkErr(err)
			fmt.Fprintf(w, "name:\t"+name)
			fmt.Fprintf(w, "\t\temail:\t"+email)
			fmt.Fprintf(w, "\t\tphone:\t"+phone)
			fmt.Fprintf(w, "\t\tdetail1:\t"+detail1)
			fmt.Fprintf(w, "\t\tdetail2:\t"+detail2)
			fmt.Fprintf(w, "\n")
		}
	} else {
		fmt.Fprintf(w, "User not exist!")
	}

}

func showData(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	vars := mux.Vars(r)

	if vars["person"] != "admin" {
		fmt.Fprintf(w, "You can't search buddy!")
		return
	}

	var retrieve *sql.Rows
	var err error

	if vars["type"] == "phone" {
		retrieve, err = db.Query("SELECT * FROM " + tbName + " WHERE phone = '" + vars["value"] + "'")
	} else if vars["type"] == "email" {
		retrieve, err = db.Query("SELECT * FROM " + tbName + " WHERE email = '" + vars["value"] + "'")
	} else {
		fmt.Fprintf(w, "what are you trying to search!")
	}
	checkErr(err)
	showAdmin(w, retrieve)

	defer retrieve.Close()
}
