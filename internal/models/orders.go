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

func (m *OrdersModel) Insert(order *Order) error {
	query := "INSERT INTO orders (created_at, status, client_id) VALUES ($1,$2,$3) RETURNING id;"
	args := []any{order.CreatedAt, order.Status, order.ClientID}

	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = tx.QueryRowContext(ctx, query, args...).Scan(&order.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, orderProduct := range order.Products {
		_, err := tx.ExecContext(ctx, "INSERT INTO orders_products (order_id,product_id,unit_price,quantity) VALUES ($1, $2, $3, $4);", order.ID, orderProduct.ProductID, orderProduct.UnitPrice, orderProduct.Quantity)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (m *OrdersModel) Get(id int64) (*Order, error) {
	if id < 1 {
		return nil, ErrNoRecord
	}
	tx, err := m.DB.Begin()
	if err != nil {
		return nil, err
	}

	query := "SELECT o.id, o.created_at, o.status, o.client_id, op.order_id, op.product_id, op.unit_priece, op.quantity FROM orders o LEFT JOIN orders_products op ON o.id = op.order_id WHERE o.id = $1;"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer result.Close()

	var order Order
	flag := true
	for result.Next() {
		flag = false
		orderProduct := new(OrderProducts)
		err := result.Scan(&order.ID, &order.CreatedAt, &order.Status, &order.ClientID, &orderProduct.OrderID, &orderProduct.ProductID, &orderProduct.UnitPrice, &orderProduct.Quantity)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		order.Products = append(order.Products, orderProduct)
	}

	if flag {
		return nil, ErrNoRecord
	}
	return &order, tx.Commit()
}

func (m *OrdersModel) Update(order *Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "UPDATE orders SET client_id = $1", order.ClientID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, "DETE FROM orders_products WHERE order_id= $1", order.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, products := range order.Products {
		_, err = tx.ExecContext(ctx, "INSERT INTO orders_products (order_id,product_id,unit_price,quantity) VALUES ($1,$2,$3,$4);", order.ID, products.ProductID, products.UnitPrice, products.Quantity)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (m *OrdersModel) Delete(id int64) error {
	if id < 1 {
		return ErrNoRecord
	}
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}
	query := "DELETE FROM orders WHERE id = $1;"

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
