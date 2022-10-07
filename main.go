package main

import (
	"fmt"
	"log"
	"net"

	"github.com/izaakdale/service-delivery/service"
	"github.com/izaakdale/utils-go/logger"
)

func main() {
	// TODO remove hard coding
	var addr = "localhost:8081"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	logger.Info(fmt.Sprintf("Listening on %s... 🦻🏻", addr))
	log.Fatal(service.New().Serve(lis))
}
