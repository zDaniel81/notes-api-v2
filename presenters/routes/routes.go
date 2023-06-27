package routes

import (
	"database/sql"
	"net/http"
	"notes-api/v2/presenters/handlers"

	"github.com/julienschmidt/httprouter"
)

func Build(router httprouter.Router, db *sql.DB) {

	router.GET("/note/:note_id", handlers.GetNote(db))
	router.GET("/notes", handlers.GetAllNotes(db))
	router.DELETE("/note/:note_id", handlers.DeleteNote(db))
	router.POST("/note/:title/:content", handlers.CreateNote(db))
	router.PUT("/note/:note_id/content/:content", handlers.UpdateNoteContent(db))
	router.PUT("/note/:note_id/title/:title", handlers.UpdateNoteTitle(db))

	if err := http.ListenAndServe(":8080", &router); err != nil {
		panic(err)
	}
}
