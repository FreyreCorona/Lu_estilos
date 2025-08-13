package models

import (
	"context"
	"database/sql"
	"time"
)

type Product struct {
	ID           int64           `json:"id"`
	Name         string          `json:"name"`
	Description  *string         `json:"description"`
	BarCode      *string         `json:"bar_code"`
	Category     *string         `json:"category"`
	InitialStock int32           `json:"initial_stock"`
	ActualStock  int32           `json:"actual_stock"`
	Price        float64         `json:"price"`
	DueDate      *time.Time      `json:"due_date"`
	Images       []*ProductImage `json:"images"`
}

type ProductImage struct {
	ID        *int64  `json:"id"`
	ProductID *int64  `json:"product_id"`
	URL       *string `json:"url"`
	Position  *int16  `json:"position"`
}

type ProductModel struct {
	DB *sql.DB
}

func (m *ProductModel) Insert(p *Product) error {
}

func (m *ProductModel) Get(id int64) (*Product, error) {
}

func (m *ProductModel) Update(p *Product) error {
}

func (m *ProductModel) Delete(id int64) error {
	if id < 1 {
		return ErrNoRecord
	}

	query := "DELETE FROM products WHERE id = $1"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	ra, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if ra == 0 {
		return ErrNoRecord
	}
	return nil
}
