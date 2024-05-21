package models

import (
	"database/sql"
	"time"
)

// Snippet is a data type ot hold data for a snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel is used to wrap a sql DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert adds a snippet to DB
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Insert adds a snippet to DB
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// Insert adds a snippet to DB
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}