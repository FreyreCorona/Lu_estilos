package models

import (
	"database/sql"
	"errors"
)

var ErrNoRecord = errors.New("no record founded")

type Models struct {
	Clients ClientModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		ClientModel{DB: db},
	}
}
