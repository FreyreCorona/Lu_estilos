package models

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNoRecord     = errors.New("no record founded")
	ErrDuplicateKey = errors.New("the email or cpf already exists")
)

type Models struct {
	Clients  ClientModel
	Products ProductModel
	Orders   OrdersModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		ClientModel{DB: db},
		ProductModel{DB: db},
		OrdersModel{DB: db},
	}
}

func HashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
