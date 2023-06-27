package main

import (
	"database/sql"
	"log"
	"notes-api/v2/infrastructure/connectors"
	routes "notes-api/v2/presenters/routes"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", connectors.SqlConnect())

	if err != nil {
		log.Panic(err)
	}

	router := *httprouter.New()

	routes.Build(router, db)
}
