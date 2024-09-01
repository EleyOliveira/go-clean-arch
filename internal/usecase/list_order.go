package usecase

import "github.com/EleyOliveira/go-clean-arch/internal/entity"

type OrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	Finalprice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
	orderRepository *entity.OrderRepository
}

func NewListOrderUseCase(repository *entity.OrderRepository) *ListOrderUseCase {
	return &ListOrderUseCase{orderRepository: repository}
}

func (l *ListOrderUseCase) ListOrders() ([]entity.Order, error) {
	return l.orderRepository.ListOrders()
}
