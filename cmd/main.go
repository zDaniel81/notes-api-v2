package main

import (
	"database/sql"
	"log"
	"notes-api/v2/infrastructure/connectors"
	presenters "notes-api/v2/infrastructure/presenters"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", connectors.SqlConnect())

	if err != nil {
		log.Panic(err)
	}

	router := *httprouter.New()
	//SqlRepository := repository.NewSqlRepository(db)

	presenters.Build(router, db)
}
