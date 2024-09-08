package service

import (
	"context"

	"github.com/EleyOliveira/go-clean-arch/internal/infra/grpc/pb"
	"github.com/EleyOliveira/go-clean-arch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	ListOrderUseCase usecase.ListOrderUseCase
}

func NewOrderService(listOrderUseCase usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		ListOrderUseCase: listOrderUseCase,
	}
}

func (s *OrderService) ListOrder(ctx context.Context, in *pb.Blank) (*pb.OrderList, error) {
	output, err := s.ListOrderUseCase.ListOrders()
	if err != nil {
		return nil, err
	}

	var ordersResponse []*pb.OrderResponse

	for _, order := range output {
		orderResponse := &pb.OrderResponse{
			Id:         order.ID,
			Tax:        float32(order.Tax),
			Price:      float32(order.Price),
			FinalPrice: float32(order.Finalprice),
		}
		ordersResponse = append(ordersResponse, orderResponse)
	}

	return &pb.OrderList{Orders: ordersResponse}, nil
}
