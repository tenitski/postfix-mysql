package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type SenderRepository interface {
	Get(login string) ([]string, error)
	Add(login string, sender string) error
	Remove(login string, sender string) error
}

type senderRepository struct {
	db *sqlx.DB
}

func NewSenderRepository(db *sqlx.DB) SenderRepository {
	return &senderRepository{db: db}
}

func (r senderRepository) Get(login string) ([]string, error) {
	var senders []string
	err := r.db.Select(&senders, "SELECT sender FROM sender WHERE login = ? ORDER BY sender", login)
	if err != nil {
		return nil, err
	}

	return senders, nil
}

func (r senderRepository) Add(login string, sender string) error {
	_, err := r.db.Exec("INSERT INTO sender (login, sender) VALUES (?, ?)", login, sender)

	return err
}

func (r senderRepository) Remove(login string, sender string) error {
	result, err := r.db.Exec("DELETE FROM sender WHERE login = ? AND sender = ?", login, sender)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
