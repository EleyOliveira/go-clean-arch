package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	usecase := NewListOrderUseCase(dbConn)

	orders, err := usecase.ListOrders()
	if err != nil {
		panic(err)
	}

	for _, order := range orders {
		fmt.Println(order.ID, order.Price, order.Tax, order.Finalprice)
	}

	/*queries := db.New(dbConn)

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
	}*/
}
