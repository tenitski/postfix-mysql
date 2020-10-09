package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	"smtpkeeper/smtp"
)

type UserRepository interface {
	Get(email string) (*smtp.User, error)
	GetAll() ([]smtp.User, error)
	Create(user smtp.User) error
	Update(user smtp.User) error
	Delete(email string) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r userRepository) Get(email string) (*smtp.User, error) {
	var user smtp.User
	err := r.db.Get(&user, "SELECT * FROM user WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r userRepository) GetAll() ([]smtp.User, error) {
	var users []smtp.User
	err := r.db.Select(&users, "SELECT * FROM user ORDER BY email")
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r userRepository) Create(user smtp.User) error {
	_, err := r.db.Exec("INSERT INTO user (email, password) VALUES (?, ?)", user.Email, user.Password)

	return err
}

func (r userRepository) Update(user smtp.User) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Make sure user exists
	var v string
	err = r.db.Get(&v, "SELECT email FROM user WHERE email = ?", user.Email)
	if err != nil {
		return err
	}

	// Update user
	_, err = tx.Exec("UPDATE user SET password = ? WHERE email = ?", user.Password, user.Email)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r userRepository) Delete(email string) error {
	result, err := r.db.Exec("DELETE FROM user WHERE email = ?", email)
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
