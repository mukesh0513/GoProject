package dbReadCreate

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var dbDriver = "mysql"
var dbUser = "root"
var dbPass = "root"
var dbName = "firstgodbNEW"
var TbName = "logGenerate"
var startNumEncode = 2
var endNumEncode = 2
var startEmEncode = 4
var endEmEncode = 10

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
	numGoRoutines := len(user)
	var wg sync.WaitGroup

	for num := 0; num < numGoRoutines; num++ {
		wg.Add(1)
		go func(num int) {

			result := user[num : num+1][0]
			allDetails := result["details"].(map[string]interface{})

			jsonString, err := json.Marshal(allDetails)
			CheckErr(err)

			insert, err := db.Query("INSERT INTO" + " " + TbName + "(name, email, phone, details) " +
				"VALUES( '" + result["name"].(string) + "', '" + result["email"].(string) + "', '" + result["phone"].(string) + "', '" + string(jsonString) + "' )")

			CheckErr(err)

			defer wg.Done()
			defer insert.Close()

		}(num)

	}
	wg.Wait()

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

				startPhone := phone[:startNumEncode]
				endPhone := phone[len(phone)-endNumEncode:]

				startEmail := email[:startEmEncode]
				endEmail := email[len(email)-endEmEncode:]

				fmt.Fprintf(w, "name : "+name)
				fmt.Fprintf(w, "\temail : "+startEmail+strings.Repeat("x", len(email)-startNumEncode-endNumEncode)+endEmail)
				fmt.Fprintf(w, "\t\tphone : "+startPhone+strings.Repeat("x", len(phone)-startNumEncode-endNumEncode)+endPhone)
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
