package models

import "database/sql"

type Client struct {
	Base
	Email    string
	CPF      string
	Password string
	Role     string
}

type ClietModel struct {
	DB *sql.DB
}
