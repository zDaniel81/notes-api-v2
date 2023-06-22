package usecases

import (
	entities "notes-api/v2/domain/entities"
	interfaces "notes-api/v2/domain/interfaces"
)

type UpdateNoteUseCase struct {
	repository interfaces.NoteRepository
}

func NewUpdateNoteUseCase(repository interfaces.NoteRepository) *UpdateNoteUseCase {
	return &UpdateNoteUseCase{
		repository: repository,
	}
}

func (uc *CreateNoteUseCase) Update(note entities.Note, title string, content string) *entities.Note {
	note.Title = title
	note.Content = content

	return &note
}
