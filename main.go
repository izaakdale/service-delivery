package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/izaakdale/lib/listener"
	"github.com/izaakdale/service-order/schema/event"
	"github.com/izaakdale/service-order/schema/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client order.OrderServiceClient
)

func main() {
	log.Printf("starting delivery service")
	cfg, err := config.LoadDefaultConfig(context.Background(), func(o *config.LoadOptions) error {
		o.Region = "eu-west-2"
		return nil
	})
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to %s", "localhost:50051")
	}
	defer conn.Close()

	client = order.NewOrderServiceClient(conn)

	err = listener.Initialise(cfg, "http://localhost:4566/000000000000/order-queue", listener.WithEndpoint("http://localhost:4566"))
	if err != nil {
		panic(err)
	}

	var errChan = make(chan error)
	go listener.Listen(process, errChan)

	for err := range errChan {
		log.Printf("listen error: %s\n", err.Error())
	}
}

func process(msg listener.Message) error {

	var ce cloudevents.Event
	err := json.Unmarshal([]byte(msg.Message), &ce)
	if err != nil {
		return err
	}

	switch ce.Type() {
	case event.TypeOrderCreated:

		var e event.OrderCreatedPayload
		err = json.Unmarshal(ce.Data(), &e)
		if err != nil {
			return err
		}
		log.Printf("starting delivery of order %s", e.ID)

		order, err := client.GetOrder(context.Background(), &order.OrderRequest{Id: e.ID})
		if err != nil {
			return err
		}
		// do something with order
		log.Printf("order %+v", order.DeliveryAddress.Name)
	case event.TypeOrderUpdated:
		log.Printf("updated hit\n")
	default:
		return fmt.Errorf("recieved an event that is not for this service: %s", ce.Type())
	}

	return nil
}
