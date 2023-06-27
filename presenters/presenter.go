package presenters

import "notes-api/v2/domain/entities"

func PresentNote(note entities.Note) *entities.Note {
	return &entities.Note{
		ID:      note.ID,
		Title:   note.Title,
		Content: note.Content,
	}
}

func PresentNotes(notes []*entities.Note) []*entities.Note {

	var result []*entities.Note

	for note := range notes {
		note := &entities.Note{
			ID:      notes[note].ID,
			Content: notes[note].Content,
			Title:   notes[note].Title,
		}
		result = append(result, note)
	}

	return result
}
