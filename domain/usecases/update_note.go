package usecases

import (
	entities "notes-api/v2/domain/entities"
	repositories "notes-api/v2/domain/repositories"
)

type UpdateNoteUseCase struct {
	repository repositories.NoteRepository
}

func NewUpdateNoteUseCase(repository repositories.NoteRepository) *UpdateNoteUseCase {
	return &UpdateNoteUseCase{
		repository: repository,
	}
}

func (uc *UpdateNoteUseCase) Call(id *string, title string, content string) (*entities.Note, error) {
	note, err := uc.repository.Update(id, title, content)

	if err != nil {
		return nil, err
	}

	return note, nil
}
