package main

import (
	"context"
	"database/sql"

	"github.com/EleyOliveira/go-clean-arch/internal/db"
	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	err = queries.CreateOrder(ctx, db.CreateOrderParams{
		ID:         uuid.New().String(),
		Tax:        2.5,
		Price:      5,
		Finalprice: 12.5,
	})

	if err != nil {
		panic(err)
	}

	orders, err := queries.ListOrders(ctx)
	if err != nil {
		panic(err)
	}

	for _, order := range orders {
		println(order.ID, order.Tax, order.Price, order.Finalprice)
	}
}
