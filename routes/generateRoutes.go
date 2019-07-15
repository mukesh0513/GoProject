package routes

import (
	. "GoProject/dbReadCreate"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRoutes() {

	r := mux.NewRouter()
	r.HandleFunc("/{person}", AllData)
	r.HandleFunc("/{person}/{name}", PersonData)
	r.HandleFunc("/{person}/{type}/{value}", Query)

	err := http.ListenAndServe(":8080", r)
	CheckErr(err)
}
