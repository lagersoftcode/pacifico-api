package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var config = GetConfig()
var db, dbError = gorm.Open("mysql", config.DBConnString)
var router = mux.NewRouter().StrictSlash(true)

func startServer() {
	fmt.Println("Starting server on address: " + config.Address)
	log.Fatal(http.ListenAndServe(config.Address, router))
}

func main() {

	if dbError != nil {
		log.Fatal(dbError)
	} else {
		RunMigrations()
		DeclareRoutes()
		startServer()
	}
}
