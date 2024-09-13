package usecase

import "github.com/EleyOliveira/go-clean-arch/internal/entity"

type OrderListOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	Finalprice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
	orderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(repository entity.OrderRepositoryInterface) *ListOrderUseCase {
	return &ListOrderUseCase{orderRepository: repository}
}

func (l *ListOrderUseCase) ListOrders() ([]OrderListOutputDTO, error) {
	orders, err := l.orderRepository.ListOrders()
	if err != nil {
		return []OrderListOutputDTO{}, err
	}

	ordersDTO := []OrderListOutputDTO{}

	for _, order := range orders {
		ordersDTO = append(ordersDTO,
			OrderListOutputDTO{ID: order.ID,
				Price:      order.Price,
				Tax:        order.Tax,
				Finalprice: order.Finalprice})
	}

	return ordersDTO, nil
}
