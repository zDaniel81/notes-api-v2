package usecases

import (
	entities "notes-api/v2/domain/entities"
	interfaces "notes-api/v2/domain/interfaces"
)

type CreateNoteUseCase struct {
	repository interfaces.NoteRepository
}

func NewCreateNoteUseCase(repository interfaces.NoteRepository) *CreateNoteUseCase {
	return &CreateNoteUseCase{
		repository: repository,
	}
}

func (uc *CreateNoteUseCase) Create(id *string, title string, content string) *entities.Note {
	return &entities.Note{
		ID:      id,
		Title:   title,
		Content: content,
	}
}
