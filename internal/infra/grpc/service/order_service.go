package service

import (
	"context"

	"github.com/EleyOliveira/go-clean-arch/internal/infra/grpc/pb"
	"github.com/EleyOliveira/go-clean-arch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	ListOrderUseCase   usecase.ListOrderUseCase
	CreateOrderUseCase usecase.CreateOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrderUseCase usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		ListOrderUseCase:   listOrderUseCase,
		CreateOrderUseCase: createOrderUseCase,
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

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.OrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}
