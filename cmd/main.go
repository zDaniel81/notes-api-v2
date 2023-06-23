package main

import (
	"database/sql"
	"fmt"
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

	var titolo string = "Titolo"
	var contenuto string = "Contenuto"

	note, err := SqlRepository.Create(&titolo, &contenuto)

	if err != nil {
		log.Panic(err)
	}

	if note == nil {
		fmt.Print("note is nil")
	}

}
