package dbReadCreate

import (
	. "GoProject/jsonStruct"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var dbDriver = "mysql"
var dbUser = "root"
var dbPass = "root"
var dbName = "firstgodb"
var TbName = "logGenerate"

func DbConn() (db *sql.DB) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func UploadData(user []Users) {

	db := DbConn()
	for i := 0; i < len(user); i++ {

		insert, err := db.Query("INSERT INTO" + " " + TbName + "(name, email, phone, detail1, detail2) " +
			"VALUES( '" + user[i].Name + "', '" + user[i].Email + "', '" + user[i].Phone + "', '" + user[i].Detail1 + "', '" +
			user[i].Detail2 + "' )")

		CheckErr(err)
		defer insert.Close()
	}
}

func PersonData(w http.ResponseWriter, r *http.Request) {

	db := DbConn()
	vars := mux.Vars(r)

	retrieve, err := db.Query("SELECT * FROM " + TbName + " WHERE name = '" + vars["name"] + "'")
	CheckErr(err)

	if vars["person"] == "user" || vars["person"] == "admin" {
		authData(w, retrieve, vars["person"])
	} else {
		panic(err.Error())
	}
	defer retrieve.Close()
}

func AllData(w http.ResponseWriter, r *http.Request) {

	db := DbConn()
	vars := mux.Vars(r)

	retrieve, err := db.Query("SELECT * FROM " + TbName)
	CheckErr(err)

	if vars["person"] == "user" || vars["person"] == "admin" {
		authData(w, retrieve, vars["person"])
	} else {
		panic(err.Error())
	}
	defer retrieve.Close()
}

func Query(w http.ResponseWriter, r *http.Request) {

	db := DbConn()
	vars := mux.Vars(r)

	if vars["person"] != "admin" {
		fmt.Fprintf(w, "You can't search buddy!")
		return
	}

	var retrieve *sql.Rows
	var err error

	if vars["type"] == "phone" {
		retrieve, err = db.Query("SELECT * FROM " + TbName + " WHERE phone = '" + vars["value"] + "'")
	} else if vars["type"] == "email" {
		retrieve, err = db.Query("SELECT * FROM " + TbName + " WHERE email = '" + vars["value"] + "'")
	} else {
		fmt.Fprintf(w, "what are you trying to search!")
	}
	CheckErr(err)
	authData(w, retrieve, "admin")

	defer retrieve.Close()
}

func authData(w http.ResponseWriter, retrieve *sql.Rows, userType string) {

	if retrieve.Next() {
		for retrieve.Next() {
			var uid int
			var name string
			var email string
			var phone string
			var detail1 string
			var detail2 string
			err := retrieve.Scan(&uid, &name, &email, &phone, &detail1, &detail2)

			CheckErr(err)
			if userType == "user" {
				fmt.Fprintf(w, "inside user")
				fmt.Fprintf(w, "name:\t"+name)
				fmt.Fprintf(w, "\t\temail:\t"+"xxxxxxxxxxxxxx")
				fmt.Fprintf(w, "\t\tphone:\t"+"xxxxxxxxxxxxxx")
				fmt.Fprintf(w, "\t\tdetail1:\t"+detail1)
				fmt.Fprintf(w, "\t\tdetail2:\t"+detail2)
				fmt.Fprintf(w, "\n")
			} else {
				fmt.Fprintf(w, "inside admin")
				fmt.Fprintf(w, "name:\t"+name)
				fmt.Fprintf(w, "\t\temail:\t"+email)
				fmt.Fprintf(w, "\t\tphone:\t"+phone)
				fmt.Fprintf(w, "\t\tdetail1:\t"+detail1)
				fmt.Fprintf(w, "\t\tdetail2:\t"+detail2)
				fmt.Fprintf(w, "\n")
			}

			fmt.Fprintf(w, "\n")
		}
	} else {
		fmt.Fprintf(w, "User not exist!")
	}

}
