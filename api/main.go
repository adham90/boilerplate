package main

import (
	"net/http"

	"github.com/adham90/boilerplate/api/handler"
	"github.com/adham90/boilerplate/pkg/utils"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	//  TODO: read from database.yml file <09-06-18, adham> //
	dbconfig := database.Config{
		Username:     "postgres",
		Password:     "root",
		Host:         "localhost",
		Port:         "5432",
		DatabaseName: "geekhubdb",
		LogMode:      true, // enable logging only in development env
	}

	db, err := database.New(dbconfig)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	db.Migrate()

	handler.MakeUserHandlers(r, db.DB)

	http.ListenAndServe(":8080", r)
}
