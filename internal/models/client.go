// Package models for hold the data types representation on the DB
package models

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Client struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	CPF      string `json:"cpf"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type ClientModel struct {
	DB *sql.DB
}

func (m *ClientModel) Insert(client *Client) error {
	query := "INSERT INTO clients  (id,name,email,cpf,password,role) VALUES ($1,$2,$3,$4,$5,$6)"
	args := []any{client.ID, client.Name, client.Email, client.CPF, client.Password, client.Role}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return err
	}
	client.ID = id
	return nil
}

func (m *ClientModel) Get(id int64) (*Client, error) {
	if id < 1 {
		return nil, ErrNoRecord
	}

	query := "SELECT id,name,email,cpf,passowrd,role FROM clients WHERE id = $1"
	var client Client

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&client.ID,
		&client.Name,
		&client.Email,
		&client.CPF,
		&client.Password,
		&client.Role)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoRecord
		default:
			return nil, err
		}
	}
	return &client, nil
}
