package main

import (
	"database/sql"
	"log"
	"notes-api/v2/infrastructure/connectors"

	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", connectors.SqlConnect())

	if err != nil {
		log.Panic(err)
	}

	//SqlRepository := repository.NewSqlRepository(db)

	//var titolo string = "Titolo"
	//var contenuto string = "Contenuto"

	if err != nil {
		log.Panic(err)
	}

}
