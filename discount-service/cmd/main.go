package main

import (
	"discount-service/core/handler"
	"discount-service/core/infra/listener"
	"discount-service/core/infra/mongodb"
	"discount-service/core/repository"
	"discount-service/core/service"
	"google.golang.org/grpc"
	"log"
)

var (
	mongodbUri      = "mongodb://localhost:27017/"
	grpcPort   uint = 50051
)

func main() {
	dbClient := mongodb.Connect(mongodbUri)

	usersCollection := dbClient.Database("discount").Collection("users")
	productsCollection := dbClient.Database("discount").Collection("products")
	userRepository := repository.NewUserRepository(usersCollection)
	productRepository := repository.NewProductRepository(productsCollection)
	discountService := service.NewService(userRepository, productRepository)
	discountGrpcHandler := handler.NewDiscountGrpcHandler(discountService)

	netListener := listener.GetNetListener(grpcPort)

	grpcServer := grpc.NewServer()

	handler.RegisterDiscountServer(grpcServer, discountGrpcHandler)

	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatal(err)
	}
}
