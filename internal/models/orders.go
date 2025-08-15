package models

import (
	"database/sql"
	"time"
)

type Order struct {
	ID        int64
	CreatedAt time.Time
	Status    string
	ClientID  int64
	Products  []*OrderProducts
}

type OrderProducts struct {
	OrderID   int64
	ProductID int64
	UnitPrice float64
	Quantity  int
}

type OrdersModel struct {
	DB *sql.DB
}

func (m *OrdersModel) Insert(order *Order) (*Order, error) {
}

func (m *OrdersModel) Get(id int64) error {
}

func (m *OrdersModel) Update(order *Order) error {
}

func (m *OrdersModel) Delete(id int64) error {
}
