package usecases

import (
	entities "notes-api/v2/domain/entities"
	repositories "notes-api/v2/domain/repositories"
)

type UpdateTitleNoteUseCase struct {
	repository repositories.NoteRepository
}

func NewUpdateTitleNoteUseCase(repository repositories.NoteRepository) *UpdateTitleNoteUseCase {
	return &UpdateTitleNoteUseCase{
		repository: repository,
	}
}

func (uc *UpdateTitleNoteUseCase) Call(id int, title *string) (*entities.Note, error) {
	note, err := uc.repository.UpdateTitle(id, title)

	if err != nil {
		return nil, err
	}

	return note, nil
}
