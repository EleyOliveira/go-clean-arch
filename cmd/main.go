package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/EleyOliveira/go-clean-arch/configs"
	"github.com/EleyOliveira/go-clean-arch/internal/infra/graphql/graph"
	"github.com/EleyOliveira/go-clean-arch/internal/infra/grpc/pb"
	"github.com/EleyOliveira/go-clean-arch/internal/infra/grpc/service"
	"github.com/EleyOliveira/go-clean-arch/internal/infra/web/webserver"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	//ctx := context.Background()

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	dbConn, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword,
		configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	listOrderUseCase := NewListOrderUseCase(dbConn)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(dbConn)
	webserver.AddHandler("/list", webOrderHandler.List)
	fmt.Println("web server inicializado na porta", configs.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	listOrderService := service.NewOrderService(*listOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, listOrderService)
	reflection.Register(grpcServer)
	fmt.Println("Servidor gRPC inicializado na porta ", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		ListOrderUseCase: *listOrderUseCase,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	fmt.Println("Inicializado servidor GraphQL na porta ", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)

	/*orders, err := usecase.ListOrders()
	if err != nil {
		panic(err)
	}

	for _, order := range orders {
		fmt.Println(order.ID, order.Price, order.Tax, order.Finalprice)
	}

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
	}*/
}
