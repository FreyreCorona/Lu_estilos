package models

import (
	"context"
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
	if id < 1 {
		return ErrNoRecord
	}
	query := "DELETE FROM orders WHERE id = $1"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.second)
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
