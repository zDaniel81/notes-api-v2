package handlers

import (
	"encoding/json"
	"net/http"
	"notes-api/v2/domain/usecases"
	repository "notes-api/v2/infrastructure/repositories"
	presenters "notes-api/v2/presenters"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetNote() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		var repository repository.SqlRepository
		uc := usecases.NewGetNoteUseCase(&repository)

		id, err := strconv.Atoi(ps.ByName("note_id"))

		if err != nil {
			//present error
		}

		out, err := uc.Call(&id)

		if err != nil {
			//present error
		}

		json.NewEncoder(w).Encode(presenters.PresentNote(*out))

	}
}
