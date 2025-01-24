package history

import (
	"database/sql"
	"fmt"
	"time"
)

type HistoryEntry struct {
	Position  int
	Question  string
	Response  string
	Timestamp time.Time
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) (*Store, error) {
	if err := initializeDB(db); err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func initializeDB(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		position INTEGER NOT NULL,
		question TEXT NOT NULL,
		response TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_position ON history(position)`)
	return err
}

func (s *Store) Save(question, response string) error {
	_, err := s.db.Exec(`INSERT INTO history (position, question, response) 
		VALUES ((SELECT COALESCE(MAX(position), 0) + 1 FROM history), ?, ?)`,
		question, response)
	return err
}

func (s *Store) List() ([]HistoryEntry, error) {
	rows, err := s.db.Query(`SELECT position, question, timestamp 
		FROM history ORDER BY position ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []HistoryEntry
	for rows.Next() {
		var entry HistoryEntry
		if err := rows.Scan(&entry.Position, &entry.Question, &entry.Timestamp); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

func (s *Store) Get(position int) (*HistoryEntry, error) {
	var entry HistoryEntry
	err := s.db.QueryRow(`SELECT position, question, response, timestamp 
		FROM history WHERE position = ?`, position).Scan(
		&entry.Position, &entry.Question, &entry.Response, &entry.Timestamp)
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (s *Store) GetLast() (*HistoryEntry, error) {
	var entry HistoryEntry
	err := s.db.QueryRow(`SELECT position, question, response, timestamp 
		FROM history ORDER BY position DESC LIMIT 1`).Scan(
		&entry.Position, &entry.Question, &entry.Response, &entry.Timestamp)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (s *Store) DeleteAll() error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.Exec("DELETE FROM history"); err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Store) Delete(position int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	res, err := tx.Exec("DELETE FROM history WHERE position = ?", position)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no entry found at position %d", position)
	}

	_, err = tx.Exec("UPDATE history SET position = position - 1 WHERE position > ?", position)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Store) Close() error {
	return s.db.Close()
}
