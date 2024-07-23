package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "grpc_connector/gen/go"
)

const (
	address = "localhost:50051"
)

func main() {
	// Устанавливаем соединение с gRPC сервером
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	client := pb.NewGreeterClient(conn)

	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Отправляем запрос на сервер
	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: "George"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
