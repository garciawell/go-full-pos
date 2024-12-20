package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/garciawell/go-full-pos/00-challenge-clean-arch/configs"
	"github.com/garciawell/go-full-pos/00-challenge-clean-arch/internal/event/handler"
	"github.com/garciawell/go-full-pos/00-challenge-clean-arch/internal/infra/graph"
	"github.com/garciawell/go-full-pos/00-challenge-clean-arch/internal/infra/grpc/pb"
	"github.com/garciawell/go-full-pos/00-challenge-clean-arch/internal/infra/grpc/service"
	"github.com/garciawell/go-full-pos/00-challenge-clean-arch/internal/infra/web/webserver"
	"github.com/garciawell/go-full-pos/00-challenge-clean-arch/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Connected to database", configs.DBHost)
	fmt.Println("Connected to RABBIT", configs.RabbitHost)

	// RABBMITMQ
	rabbitMQChannel := getRabbitMQChannel(configs)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	// WEB SERVER
	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)

	webserver.AddHandler("/order", webOrderHandler.Create)
	webserver.AddHandler("/orders", webOrderHandler.List)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	// GRPC SERVER
	grpcServer := grpc.NewServer()
	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	listOrderUseCase := NewListOrderUseCase(db)
	createOrderService := service.NewOrdersService(*createOrderUseCase, *listOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	// GRAPHQL SERVER
	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrderUseCase:   *listOrderUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

func getRabbitMQChannel(env *configs.Conf) *amqp.Channel {
	// fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cnonfigs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName)
	// "amqp://guest:guest@rabbitmq:5672/"
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", "guest", "guest", env.RabbitHost, "5672"))
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
