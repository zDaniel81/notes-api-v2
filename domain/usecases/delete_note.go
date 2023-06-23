package usecases

import (
	"notes-api/v2/domain/entities"
	repositories "notes-api/v2/domain/repositories"
)

type DeleteNoteUseCase struct {
	repository repositories.NoteRepository
}

func NewDeleteUseCase(repository repositories.NoteRepository) *DeleteNoteUseCase {
	return &DeleteNoteUseCase{
		repository: repository,
	}
}

func (uc *DeleteNoteUseCase) Call(id *int) (*entities.Note, error) {
	note, err := uc.repository.Delete(id)

	if err != nil {
		return nil, err
	}

	return note, err
}
