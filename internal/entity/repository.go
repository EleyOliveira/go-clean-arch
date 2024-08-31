package entity

import "database/sql"

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (o *OrderRepository) ListOrders() ([]Order, error) {

	orders := []Order{{ID: "20", Tax: 1.5, Price: 56.9},
		{ID: "22", Tax: 5.6, Price: 569}}

	return orders, nil
}
