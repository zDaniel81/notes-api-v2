package repositories

import entities "notes-api/v2/domain/entities"

type NoteRepository interface {
	Create(id *string, title string, content string) (*entities.Note, error)
	Update(id *string, title string, content string) (*entities.Note, error)
	Delete(id *string) (*entities.Note, error)
	GetById(id *string) (*entities.Note, error)
	GetAll() ([]*entities.Note, error)
}
