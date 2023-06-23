package usecases

import (
	"notes-api/v2/domain/entities"
	repositories "notes-api/v2/domain/repositories"
)

type GetNoteUseCase struct {
	repository repositories.NoteRepository
}

func NewGetNoteUseCase(repository repositories.NoteRepository) *GetNoteUseCase {
	return &GetNoteUseCase{
		repository: repository,
	}
}

func (uc *GetNoteUseCase) Call(id *int64) (*entities.Note, error) {
	note, err := uc.repository.GetById(id)

	if err != nil {
		return nil, err
	}

	return note, err
}
