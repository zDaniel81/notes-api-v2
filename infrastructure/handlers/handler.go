package handlers

import (
	"encoding/json"
	"net/http"
	"notes-api/v2/domain/usecases"
	presenters "notes-api/v2/infrastructure/presenters/notes"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetNote(uc *usecases.GetNoteUseCase) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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
