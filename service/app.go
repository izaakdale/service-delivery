package service

import (
	"context"

	"github.com/izaakdale/service-delivery/dao"
	"github.com/izaakdale/service-delivery/model/delivery"
	"github.com/izaakdale/utils-go/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	delivery.DeliveryServiceServer
}

func New() *grpc.Server {
	dao.Init("ordering-app", "eu-west-2")
	gs := grpc.NewServer()
	delivery.RegisterDeliveryServiceServer(gs, &Server{})
	reflection.Register(gs)
	return gs
}

func (s *Server) Delivery(ctx context.Context, in *delivery.DeliveryRequest) (*delivery.DeliveryResponse, error) {

	logger.Info("Order ID: " + in.OrderId)

	dao.UpdateStatus(&dao.DeliveryRecord{
		OrderId:        "ORDER#" + in.OrderId,
		DeliveryStatus: delivery.ORDER_STATUS_name[int32(delivery.ORDER_STATUS_PROCESSING)],
		Address:        in.Address,
	})

	return &delivery.DeliveryResponse{
		Status: delivery.ORDER_STATUS_PROCESSING,
	}, nil
}
