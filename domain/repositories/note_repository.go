package repositories

import entities "notes-api/v2/domain/entities"

type NoteRepository interface {
	Create(title string, content string) (*entities.Note, error)
	Update(id *int64, title *string, content string) (*entities.Note, error)
	Delete(id *int64) (*entities.Note, error)
	GetById(id *int64) (*entities.Note, error)
	GetAll() ([]*entities.Note, error)
}
