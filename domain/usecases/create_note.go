package usecases

import (
	entities "notes-api/v2/domain/entities"
	repositories "notes-api/v2/domain/repositories"
)

type CreateNoteUseCase struct {
	repository repositories.NoteRepository
}

func NewCreateNoteUseCase(repository repositories.NoteRepository) *CreateNoteUseCase {
	return &CreateNoteUseCase{
		repository: repository,
	}
}

func (uc *CreateNoteUseCase) Call(id *string, title string, content string) (*entities.Note, error) {
	note, err := uc.repository.Create(id, title, content)

	if err != nil {
		return nil, err
	}

	return note, nil
}
