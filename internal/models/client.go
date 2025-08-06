package models

import "database/sql"

type Client struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	CPF      string `json:"cpf"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type ClietModel struct {
	DB *sql.DB
}
