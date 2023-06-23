package usecases

import (
	"notes-api/v2/domain/entities"
	repositories "notes-api/v2/domain/repositories"
)

type GetNotesUseCase struct {
	repository repositories.NoteRepository
}

func NewGetNotesUseCase(repository repositories.NoteRepository) *GetNotesUseCase {
	return &GetNotesUseCase{
		repository: repository,
	}
}

func (uc *GetNotesUseCase) Call() ([]*entities.Note, error) {
	notes, err := uc.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return notes, err
}
