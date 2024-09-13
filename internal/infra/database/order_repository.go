package database

import (
	"context"
	"database/sql"

	"github.com/EleyOliveira/go-clean-arch/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (o *OrderRepository) ListOrders() ([]entity.Order, error) {

	ctx := context.Background()
	queries := entity.New(o.Db)

	orders, err := queries.ListOrders(ctx)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}
