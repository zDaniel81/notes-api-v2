package infrastructure

import (
	"database/sql"
	entities "notes-api/v2/domain/entities"
	repositories "notes-api/v2/domain/repositories"
)

type SqlRepository struct {
	db *sql.DB
}

func NewSqlRepository(db *sql.DB) *SqlRepository {
	return &SqlRepository{
		db: db,
	}
}

func (*SqlRepository) Create(id *string, title string, content string) (*entities.Note, error) {
	panic("unimplemented")
}

func (*SqlRepository) Delete(id *string) (*entities.Note, error) {
	panic("unimplemented")
}

func (*SqlRepository) GetAll() ([]*entities.Note, error) {
	panic("unimplemented")
}

func (*SqlRepository) GetById(id *string) (*entities.Note, error) {
	panic("unimplemented")
}

func (*SqlRepository) Update(id *string, title string, content string) (*entities.Note, error) {
	panic("unimplemented")
}

var _ repositories.NoteRepository = &SqlRepository{
	db: &sql.DB{},
}
