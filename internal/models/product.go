package models

import (
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

func (p *ProductModel) Insert(p *Product) error {
}

func (p *ProductModel) Get(id int64) (*Product, error) {
}

func (p *ProductModel) Update(p *Product) error {
}

func (p *ProductModel) Delete(id int64) error {
}
