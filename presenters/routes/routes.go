package routes

import (
	"database/sql"
	"net/http"
	"notes-api/v2/presenters/handlers"

	"github.com/julienschmidt/httprouter"
)

func Build(router httprouter.Router, db *sql.DB) {

	router.GET("/note/:note_id", handlers.GetNote())

	if err := http.ListenAndServe(":8080", &router); err != nil {
		panic(err)
	}
}
