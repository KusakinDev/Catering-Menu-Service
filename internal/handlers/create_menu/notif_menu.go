package createmenu

import (
	"context"
	"time"

	pb "github.com/KusakinDev/Catering-Menu-Service/internal/services/notif_new_menu/notif_new_menu"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NotifNewMenu(mess string) string {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logrus.Errorln("Error connect:", err)
		return ""
	}
	defer conn.Close()

	client := pb.NewNotifNewMenuServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.MenuRequest{Message: mess}

	res, err := client.NotifNewMenu(ctx, req)
	if err != nil {
		logrus.Errorln("Error send message:", err)
		return ""
	}

	return res.Message
}
