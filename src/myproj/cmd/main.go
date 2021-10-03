package main

import (
	"flag"

	"log"
	"my-project-mux/main/internal/database"
	"my-project-mux/main/internal/repository"
	"my-project-mux/main/internal/service"

	"net/http"

	"github.com/gorilla/mux"
)


func main() {

	//init db
	var port, dbhost, dbname, dbusername, dbpassword, disableTLS string
	var dbport int
	flag.StringVar(&port, "port", "3000", "port for open service")
	flag.StringVar(&dbhost, "dbhost", "localhost", "database host name")
	flag.IntVar(&dbport, "dbport", 5432, "database port")
	flag.StringVar(&dbname, "dbname", "mydb", "database schema name")
	flag.StringVar(&dbusername, "dbusername", "postgres", "database user name")
	flag.StringVar(&dbpassword, "dbpassword", "1234", "database password")
	flag.StringVar(&disableTLS, "disableTLS", "Y", "database disableTLS[Y/n]")
	flag.Parse()
	var databaseTSL bool
	if disableTLS == "n" {
		databaseTSL = false
	} else {
		databaseTSL = true
	}
	dbConfig := database.Config{
		User:       dbusername,
		Password:   dbpassword,
		Host:       dbhost,
		Port:       dbport,
		Name:       dbname,
		DisableTLS: databaseTSL,
	}

	db, err := database.Open(dbConfig)
	if err != nil {
		log.Fatal("connecting database fail", err)
	}
	repository.InitDB(db)

	//service 
	r := mux.NewRouter()
	service.InitRoute(r)

	// r.HandleFunc("/api/products", product.GetProduct).Methods("GET")
	// r.HandleFunc("/api/products", product.PostProduct).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080",r))
}