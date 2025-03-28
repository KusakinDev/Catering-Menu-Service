/*
 * Catering service
 *
 * Menu service.
 *
 * API version: 1.0.0
 */

package main

import (
	"log"
	"net"

	pb "github.com/KusakinDev/Catering-Menu-Service/internal/services/get_dish/get_dish_serv"
	"github.com/KusakinDev/Catering-Menu-Service/internal/services/get_dish/server"

	sw "github.com/KusakinDev/Catering-Menu-Service/internal/routes"
	"google.golang.org/grpc"
)

func main() {
	// REST
	go func() {
		routes := sw.ApiHandleFunctions{}
		log.Printf("REST server started on port :8080")
		router := sw.NewRouter(routes)
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to run REST server: %v", err)
		}
	}()

	// gRPC
	go func() {
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterDishServiceServer(grpcServer, &server.Server{})

		log.Println("gRPC server is running on port :50051")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	select {}
}
