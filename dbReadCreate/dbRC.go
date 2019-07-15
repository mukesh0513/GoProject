package dbReadCreate

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var dbDriver = "mysql"
var dbUser = "root"
var dbPass = "root"
var dbName = "firstgodbNEW"
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

func UploadData(user []map[string]interface{}) {

	db := DbConn()

	for _, result := range user {
		allDetails := result["details"].(map[string]interface{})

		jsonString, err := json.Marshal(allDetails)
		CheckErr(err)

		insert, err := db.Query("INSERT INTO" + " " + TbName + "(name, email, phone, details) " +
			"VALUES( '" + result["name"].(string) + "', '" + result["email"].(string) + "', '" + result["phone"].(string) + "', '" + string(jsonString) + "' )")

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
			var details string
			err := retrieve.Scan(&uid, &name, &email, &phone, &details)

			CheckErr(err)
			if userType == "user" {
				fmt.Fprintf(w, "name : "+name)
				fmt.Fprintf(w, "\temail : "+"xxxxxxxxxxxxxx")
				fmt.Fprintf(w, "\t\tphone : "+"xxxxxxxxxxxxxx")
				fmt.Fprintf(w, "\t\tdetail1 : "+details)
				fmt.Fprintf(w, "\n")
			} else {
				fmt.Fprintf(w, "name : "+name)
				fmt.Fprintf(w, "\temail : "+email)
				fmt.Fprintf(w, "\t\tphone : "+phone)
				fmt.Fprintf(w, "\t\tdetail1 : "+details)
				fmt.Fprintf(w, "\n")
			}

			fmt.Fprintf(w, "\n")
		}
	} else {
		fmt.Fprintf(w, "User not exist!")
	}

}
