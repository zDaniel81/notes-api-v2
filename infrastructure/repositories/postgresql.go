package repository

import (
	"database/sql"
	"errors"
	entities "notes-api/v2/domain/entities"
	"notes-api/v2/domain/repositories"
)

type SqlRepository struct {
	db *sql.DB
}

func NewSqlRepository(db *sql.DB) *SqlRepository {
	return &SqlRepository{
		db: db,
	}
}

var _ repositories.NoteRepository = &SqlRepository{
	db: &sql.DB{},
}

func (repo *SqlRepository) Create(title *string, content *string) (*entities.Note, error) {

	sqlStatement := `
	INSERT INTO notes(title, content) VALUES ($1, $2) RETURNING id`

	var note entities.Note

	err := repo.db.QueryRow(sqlStatement, title, content).Scan(&note.ID)

	if err != nil {
		return nil, err
	}

	note = entities.Note{
		ID:      note.ID,
		Title:   *title,
		Content: *content,
	}

	return &note, nil

}

func (repo *SqlRepository) Delete(id *int) (*entities.Note, error) {

	sqlStatement := `
	DELETE FROM NOTES
	WHERE id = $1`

	_, err := repo.db.Exec(sqlStatement, &id)

	if err != nil {
		return nil, err
	}

	note := entities.Note{
		ID: *id,
	}

	return &note, nil
}

func (repo *SqlRepository) GetAll() ([]*entities.Note, error) {

	sqlStatement := `SELECT * FROM notes;`

	rows, err := repo.db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	notes := []*entities.Note{}

	for rows.Next() {

		var note entities.Note
		err = rows.Scan(&note.ID, &note.Title, &note.Content)

		if err != nil {
			return nil, err
		}
		notes = append(notes, &note)
	}

	return notes, nil
}

func (repo *SqlRepository) GetById(id *int) (*entities.Note, error) {

	sqlStatement := `SELECT title, content FROM notes WHERE id=$1`

	var title string
	var content string

	err := repo.db.QueryRow(sqlStatement, id).Scan(title, content)

	if err == nil {
		return nil, errors.New("test")
	}

	note := entities.Note{
		ID:      *id,
		Title:   title,
		Content: content,
	}

	return &note, nil
}

func (repo *SqlRepository) UpdateContent(id *int, content *string) (*entities.Note, error) {

	note, err := repo.GetById(id)

	if err != nil {
		return nil, err
	}

	sqlStatement := `
		UPDATE notes
		SET
		content=$1`

	_, err = repo.db.Exec(sqlStatement, content)

	if err != nil {
		return nil, err
	}

	note.Content = *content
	return note, nil

}

func (repo *SqlRepository) UpdateTitle(id *int, title *string) (*entities.Note, error) {

	note, err := repo.GetById(id)

	if err != nil {
		return nil, err
	}

	sqlStatement := `
		UPDATE notes
		SET
		title=$1`

	_, err = repo.db.Exec(sqlStatement, title)

	if err != nil {
		return nil, err
	}

	note.Title = *title
	return note, nil

}
