package interfaces

import entities "notes-api/v2/domain/entities"

type NoteRepository interface {
	Create(id *string, title string, content string) (*entities.Note, error)
	Update(note *entities.Note, title string, content string) (*entities.Note, error)
	Delete(note *entities.Note) (*entities.Note, error)
	GetById(id string) (*entities.Note, error)
	GetAll() ([]*entities.Note, error)
}
