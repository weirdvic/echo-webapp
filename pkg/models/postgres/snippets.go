package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/weirdvic/echo-tutorial/pkg/models"
)

type SnippetModel struct {
	DB *pgxpool.Pool
}

// Insert - Метод для создания новой заметки в базе дынных.
func (m *SnippetModel) Insert(title, content, expires_at string) error {
	query := `INSERT INTO snippets (title, content, created_at, expires_at)
    VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + '$4 day'::interval))`
	_, err := m.DB.Exec(context.Background(), query, title, content, expires_at)
	if err != nil {
		return err
	}
	return nil
}

// Get - Метод для возвращения данных заметки по её идентификатору ID.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	query := `SELECT id, title, content, created_at, expires_at FROM snippets
    WHERE expires_at > CURRENT_TIMESTAMP AND id = $1`
	s := &models.Snippet{}
	err := m.DB.QueryRow(context.Background(), query, id).Scan(
		&s.ID,
		&s.Title,
		&s.Content,
		&s.CreatedAt,
		&s.ExpiresAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// Latest - Метод возвращает 10 наиболее часто используемых заметок.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	query := `SELECT id, title, content, created_at, expires_at FROM snippets
	WHERE expires_at > CURRENT_TIMESTAMP ORDER BY created_at DESC LIMIT 10`
	rows, err := m.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []*models.Snippet
	for rows.Next() {
		s := &models.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.ExpiresAt)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}
