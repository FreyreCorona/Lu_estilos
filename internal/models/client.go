// Package models for hold the data types representation on the DB
package models

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
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
	query := "INSERT INTO clients  (name,email,cpf,password,role) VALUES ($1,$2,$3,$4,$5) RETURNING id"
	args := []any{client.Name, client.Email, client.CPF, client.Password, client.Role}

	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = tx.QueryRowContext(ctx, query, args...).Scan(&client.ID)
	if err != nil {
		tx.Rollback()
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return ErrDuplicateKey
		}
		return err
	}
	return tx.Commit()
}

func (m *ClientModel) Get(id int64) (*Client, error) {
	if id < 1 {
		return nil, ErrNoRecord
	}

	query := "SELECT id,name,email,cpf,password,role FROM clients WHERE id = $1"
	var client Client

	tx, err := m.DB.Begin()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = tx.QueryRowContext(ctx, query, id).Scan(
		&client.ID,
		&client.Name,
		&client.Email,
		&client.CPF,
		&client.Password,
		&client.Role)
	if err != nil {
		tx.Rollback()
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoRecord
		default:
			return nil, err
		}
	}
	return &client, tx.Commit()
}

func (m *ClientModel) Update(client *Client) error {
	query := "UPDATE clients SET name = $1, email = $2, cpf = $3, password = $4, role = $5 WHERE id = $6"
	args := []any{client.Name, client.Email, client.CPF, client.Password, client.Role, client.ID}

	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (m *ClientModel) Delete(id int64) error {
	if id < 1 {
		return ErrNoRecord
	}

	query := "DELETE FROM clients WHERE id = $1"

	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	ra, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if ra == 0 {
		return ErrNoRecord
	}
	return tx.Commit()
}
