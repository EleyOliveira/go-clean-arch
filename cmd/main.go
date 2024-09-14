package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/EleyOliveira/go-clean-arch/configs"
	"github.com/EleyOliveira/go-clean-arch/internal/event/handler"
	"github.com/EleyOliveira/go-clean-arch/internal/infra/graphql/graph"
	"github.com/EleyOliveira/go-clean-arch/internal/infra/grpc/pb"
	"github.com/EleyOliveira/go-clean-arch/internal/infra/grpc/service"
	"github.com/EleyOliveira/go-clean-arch/internal/infra/web/webserver"
	"github.com/EleyOliveira/go-clean-arch/pkg/events"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
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

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	listOrderUseCase := NewListOrderUseCase(dbConn)
	createOrderUseCase := NewCreateOrderUseCase(dbConn, eventDispatcher)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(dbConn, eventDispatcher)
	webserver.Router.Route("/order", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Post("/", webOrderHandler.Create)
		r.Get("/", webOrderHandler.List)
	})
	fmt.Println("web server inicializado na porta", configs.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	orderService := service.NewOrderService(*createOrderUseCase, *listOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, orderService)
	reflection.Register(grpcServer)
	fmt.Println("Servidor gRPC inicializado na porta ", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		ListOrderUseCase:   *listOrderUseCase,
		CreateOrderUseCase: *createOrderUseCase,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	fmt.Println("Inicializado servidor GraphQL na porta ", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

func getRabbitMQChannel() *amqp.Channel {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", configs.RABBITMQ_USER, configs.RABBITMQ_PASSWORD,
		configs.RABBITMQ_HOST, configs.RABBITMQ_SERVER_PORT))

	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
