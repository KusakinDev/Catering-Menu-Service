package server

import (
	"context"
	"errors"

	dishmodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/dish_model"
	pb "github.com/KusakinDev/Catering-Menu-Service/internal/services/get_dish/get_dish_serv"
)

type Server struct {
	pb.UnimplementedDishServiceServer
}

func (s *Server) GetDishes(ctx context.Context, req *pb.DishRequest) (*pb.DishesResponse, error) {
	var dish dishmodel.Dish

	dishes, err := dish.GetAllFromTable()
	if err != 200 {
		return nil, errors.New("error get from db")
	}

	var dishes_gRPC []*pb.Dish

	for _, dish := range dishes {
		dish_gRPC := dish.ToGRPCDish()
		dishes_gRPC = append(dishes_gRPC, dish_gRPC)
	}

	return &pb.DishesResponse{Dishes: dishes_gRPC}, nil
}
