package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"notes-api/v2/domain/usecases"
	repository "notes-api/v2/infrastructure/repositories"
	presenters "notes-api/v2/presenters"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetNote(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		repository := repository.NewSqlRepository(db)
		uc := usecases.NewGetNoteUseCase(repository)

		id, err := strconv.Atoi(ps.ByName("note_id"))

		if err != nil {
			presenters.PresentErrors(err)
		}

		out, err := uc.Call(id)

		if err != nil {
			presenters.PresentErrors(err)
		}

		json.NewEncoder(w).Encode(presenters.PresentNote(*out))

	}
}

func GetAllNotes(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		repository := repository.NewSqlRepository(db)
		uc := usecases.NewGetNotesUseCase(repository)

		out, err := uc.Call()

		if err != nil {
			presenters.PresentErrors(err)
		}

		json.NewEncoder(w).Encode(presenters.PresentNotes(out))

	}
}

func DeleteNote(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		repository := repository.NewSqlRepository(db)
		uc := usecases.NewDeleteUseCase(repository)

		url_id := ps.ByName("note_id")

		id, err := strconv.Atoi(url_id)

		if err != nil {
			return
		}

		out, err := uc.Call(id)

		if err != nil {
			presenters.PresentErrors(err)
		}

		json.NewEncoder(w).Encode(presenters.PresentNote(*out))

	}
}

func CreateNote(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		repository := repository.NewSqlRepository(db)
		uc := usecases.NewCreateNoteUseCase(repository)

		title := ps.ByName("title")
		content := ps.ByName("content")

		out, err := uc.Call(&title, &content)

		if err != nil {
			presenters.PresentErrors(err)
		}

		json.NewEncoder(w).Encode(presenters.PresentNote(*out))

	}
}

func UpdateNoteContent(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		repository := repository.NewSqlRepository(db)
		uc := usecases.NewUpdateContentNoteUseCase(repository)

		content := ps.ByName("content")

		id, err := strconv.Atoi(ps.ByName("note_id"))

		if err != nil {
			presenters.PresentErrors(err)
		}

		out, err := uc.Call(id, &content)

		if err != nil {
			presenters.PresentErrors(err)
		}

		json.NewEncoder(w).Encode(presenters.PresentNote(*out))

	}
}

func UpdateNoteTitle(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		repository := repository.NewSqlRepository(db)
		uc := usecases.NewUpdateTitleNoteUseCase(repository)

		title := ps.ByName("title")

		id, err := strconv.Atoi(ps.ByName("note_id"))

		if err != nil {
			presenters.PresentErrors(err)
		}

		out, err := uc.Call(id, &title)

		if err != nil {
			presenters.PresentErrors(err)
		}

		json.NewEncoder(w).Encode(presenters.PresentNote(*out))

	}
}
