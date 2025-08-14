package models

import (
	"context"
	"database/sql"
	"errors"
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
		return nil, ErrNoRecord
	}
	query := "SELECT p.id,p.name,p.description,p.bar_code,p.category,p.initial_stock,p.actual_stock,p.price,p.due_date,i.id,i.product_id,i.url,i.position FROM products p LEFT JOIN product_images i ON p.id = i.product_id WHERE id = $1 ORDER BY i.position;"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoRecord
		default:
			return nil, err
		}
	}
	defer result.Close()

	var product Product
	for result.Next() {
		image := new(ProductImage)
		err := result.Scan(&product.ID, &product.Name, &product.Description, &product.BarCode, &product.Category, &product.InitialStock, &product.ActualStock, &product.Price, &product.DueDate, &image.ID, &image.ProductID, &image.URL, &image.Position)
		if err != nil {
			return nil, err
		}
		if image.ID != nil {
			product.Images = append(product.Images, image)
		}
	}

	return &product, nil
}

func (m *ProductModel) Update(product *Product) error {
	query := "UPDATE products SET name=$1, description=$2, bar_code=$3, category=$4, initial_stock=$5, actual_stock=$6, price=$7, due_date=$8 WHERE id = $9"
	args := []any{product.Name, product.Description, product.BarCode, product.Category, product.InitialStock, product.ActualStock, product.Price, product.DueDate, product.ID}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
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
