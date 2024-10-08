package entity

type OrderRepositoryInterface interface {
	ListOrders() ([]Order, error)
	Save(order *Order) error
}
