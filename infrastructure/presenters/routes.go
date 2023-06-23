package presenters

import (
	"database/sql"
	"net/http"
	"notes-api/v2/domain/usecases"
	"notes-api/v2/infrastructure/handlers"
	repository "notes-api/v2/infrastructure/repositories"

	"github.com/julienschmidt/httprouter"
)

func Build(router httprouter.Router, db *sql.DB) {

	router.GET("/note/:note_id", handlers.GetNote(usecases.NewGetNoteUseCase(repository.NewSqlRepository(db))))

	if err := http.ListenAndServe(":8081", &router); err != nil {
		panic(err)
	}
}
