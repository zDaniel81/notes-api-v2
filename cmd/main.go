package main

import (
	"database/sql"
	"log"
	"notes-api/v2/infrastructure/connectors"
	repository "notes-api/v2/infrastructure/repositories"

	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", connectors.SqlConnect())

	if err != nil {
		log.Panic(err)
	}

	SqlRepository := repository.NewSqlRepository(db)

	SqlRepository.Create("test", "content")

}
