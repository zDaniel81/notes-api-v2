package repository

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

func (repo *SqlRepository) Create(title string, content string) (*entities.Note, error) {

	sqlStatement := `
	INSERT INTO notes(title, content) VALUES ($1, $2)`

	result, err := repo.db.Exec(sqlStatement, title, content)

	if err != nil {
		return nil, err
	}

	db_id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	note := entities.Note{
		ID:      db_id,
		Title:   title,
		Content: content,
	}

	return &note, nil

}

func (repo *SqlRepository) Delete(id *int64) (*entities.Note, error) {

	sqlStatement := `
	DELETE FROM NOTES
	WHERE id = $1`

	result, err := repo.db.Exec(sqlStatement, id)

	if err != nil {
		return nil, err
	}

	db_id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	note := entities.Note{
		ID:    db_id,
		Title: "Note Deleted",
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

func (repo *SqlRepository) GetById(id *int64) (*entities.Note, error) {

	sqlStatement := `SELECT * FROM notes WHERE ID=$1`

	result, err := repo.db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	var note entities.Note

	err = result.Scan(&note.ID, &note.Title, &note.Content)

	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (repo *SqlRepository) Update(id *int64, title *string, content string) (*entities.Note, error) {

	var result sql.Result
	var err error

	if title != nil {
		sqlStatement := `
		UPDATE notes
		SET
		title=$1,
		content=$2`

		result, err = repo.db.Exec(sqlStatement, title, content)

	} else {
		sqlStatement := `
		UPDATE notes
		SET
		content=$1`

		result, err = repo.db.Exec(sqlStatement, content)

	}

	if err != nil {
		return nil, err
	}

	db_id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	note := entities.Note{
		ID:      db_id,
		Title:   *title,
		Content: content,
	}

	return &note, nil

}

var _ repositories.NoteRepository = &SqlRepository{
	db: &sql.DB{},
}
