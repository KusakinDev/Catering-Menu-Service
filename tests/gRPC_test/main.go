package main

import (
	"context"
	"log"
	"time"

	pb "github.com/KusakinDev/Catering-Menu-Service/internal/services/get_dish/get_dish_serv"

	"google.golang.org/grpc"
)

func main() {
	// Создайте контекст с тайм-аутом для управления подключением
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Подключитесь к серверу gRPC
	conn, err := grpc.DialContext(ctx, ":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDishServiceClient(conn)

	// Создание временного контекста для запроса
	reqCtx, reqCancel := context.WithTimeout(context.Background(), time.Second)
	defer reqCancel()

	// Вызов gRPC метода
	res, err := client.GetDish(reqCtx, &pb.DishRequest{Id: 1})
	if err != nil {
		log.Fatalf("Could not get dish: %v", err)
	}

	log.Printf("Dish: %v", res)
}
