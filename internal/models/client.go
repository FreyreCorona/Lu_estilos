package models

import "database/sql"

type Client struct {
	Base
	Email    string `json:"email"`
	CPF      string `json:"cpf"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type ClietModel struct {
	DB *sql.DB
}
