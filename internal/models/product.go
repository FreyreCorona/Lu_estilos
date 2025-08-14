package models

import (
	"context"
	"database/sql"
	"time"
)

type Product struct {
	ID           int64           `json:"id"`
	Name         string          `json:"name"`
	Description  *string         `json:"description,omitempty"`
	BarCode      *string         `json:"bar_code,omitempty"`
	Category     *string         `json:"category,omitempty"`
	InitialStock int32           `json:"initial_stock"`
	ActualStock  int32           `json:"actual_stock"`
	Price        float64         `json:"price"`
	DueDate      *time.Time      `json:"due_date,omitempty"`
	Images       []*ProductImage `json:"images,omitempty"`
}

type ProductImage struct {
	ID        *int64  `json:"id,omitempty"`
	ProductID *int64  `json:"product_id,omitempty"`
	URL       *string `json:"url,omitempty"`
	Position  *int16  `json:"position,omitempty"`
}

type ProductModel struct {
	DB *sql.DB
}

func (m *ProductModel) Insert(product *Product) error {
	query := "INSERT INTO products (name,description,bar_code,category,initial_stock,actual_stock,price,due_date) VALUES($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id;"
	args := []any{product.Name, product.Description, product.BarCode, product.Category, product.InitialStock, product.ActualStock, product.Price, product.DueDate}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&product.ID)
	if err != nil {
		return err
	}
	for _, image := range product.Images {
		_, err := m.DB.ExecContext(ctx, "INSTERT INTO product_images (product_id,url,position) VALUES ($1,$2,$3); ", product.ID, image.URL, image.Position)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *ProductModel) Get(id int64) (*Product, error) {
	if id < 1 {
		return ErrNoRecord
	}
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
