package repositories

import entities "notes-api/v2/domain/entities"

type NoteRepository interface {
	Create(title *string, content *string) (*entities.Note, error)
	UpdateContent(id int, content *string) (*entities.Note, error)
	UpdateTitle(id int, title *string) (*entities.Note, error)
	Delete(id int) (*entities.Note, error)
	GetById(id int) (*entities.Note, error)
	GetAll() ([]*entities.Note, error)
}
