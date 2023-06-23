package presenters

import "notes-api/v2/domain/entities"

func PresentNote(note entities.Note) *entities.Note {
	return &entities.Note{
		ID:      note.ID,
		Title:   note.Title,
		Content: note.Content,
	}
}
