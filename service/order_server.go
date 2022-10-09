package service

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/izaakdale/service-delivery/dao"
	"github.com/izaakdale/service-delivery/model/delivery"
	"github.com/izaakdale/utils-go/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	delivery.DeliveryServiceServer
}

func RunGrpcServer() {

	dao.Init("ordering-app", "eu-west-2")
	gs := grpc.NewServer()
	delivery.RegisterDeliveryServiceServer(gs, &GrpcServer{})
	reflection.Register(gs)

	var addr = "localhost:50001"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	logger.Info(fmt.Sprintf("Listening on %s... 🦻🏻", addr))
	log.Fatal(gs.Serve(lis))
}

func (s *GrpcServer) Delivery(ctx context.Context, in *delivery.DeliveryRequest) (*delivery.DeliveryResponse, error) {

	logger.Info("Order ID: " + in.OrderId)

	dao.UpdateStatus(&dao.DeliveryRecord{
		OrderId:    dao.OrderPrefixPK + in.OrderId,
		RecordType: dao.StatusSK,
		Status:     delivery.ORDER_STATUS_name[int32(delivery.ORDER_STATUS_PROCESSING)],
		Address:    in.Address,
	})

	// factory is a fake...
	// go factory.UpdateStatuses(in)

	return &delivery.DeliveryResponse{
		Status: delivery.ORDER_STATUS_name[int32(delivery.ORDER_STATUS_PROCESSING)],
	}, nil
}
