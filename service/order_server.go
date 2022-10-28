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

func RunGrpcServer(addr string) {
	gs := grpc.NewServer()
	delivery.RegisterDeliveryServiceServer(gs, &GrpcServer{})
	reflection.Register(gs)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	logger.Info(fmt.Sprintf("Listening on %s... 🦻🏻", addr))
	log.Fatal(gs.Serve(lis))
}

func (s *GrpcServer) Delivery(ctx context.Context, in *delivery.DeliveryRequest) (*delivery.DeliveryResponse, error) {

	logger.Info("Order ID: " + in.OrderId)

	dao.UpdateStatus(&dao.StatusRecord{
		OrderId:    dao.OrderPrefixPK + in.OrderId,
		RecordType: dao.StatusSK,
		Status:     delivery.ORDER_STATUS_name[int32(delivery.ORDER_STATUS_PROCESSING)],
	})

	dao.StoreAddress(&dao.AddressRecord{
		OrderId:    dao.OrderPrefixPK + in.OrderId,
		RecordType: dao.AddressSK,
		Address:    in.Address,
	})

	// factory is a fake...
	// go factory.UpdateStatuses(in)

	return &delivery.DeliveryResponse{
		Status: delivery.ORDER_STATUS_name[int32(delivery.ORDER_STATUS_PROCESSING)],
	}, nil
}
