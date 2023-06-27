package usecases

import (
	entities "notes-api/v2/domain/entities"
	repositories "notes-api/v2/domain/repositories"
)

type UpdateContentNoteUseCase struct {
	repository repositories.NoteRepository
}

func NewUpdateContentNoteUseCase(repository repositories.NoteRepository) *UpdateContentNoteUseCase {
	return &UpdateContentNoteUseCase{
		repository: repository,
	}
}

func (uc *UpdateContentNoteUseCase) Call(id int, content *string) (*entities.Note, error) {
	note, err := uc.repository.UpdateContent(id, content)

	if err != nil {
		return nil, err
	}

	return note, nil
}
