package main

import (
	"log"

	"github.com/izaakdale/service-delivery/service"
	"github.com/kelseyhightower/envconfig"
)

type Spec struct {
	Address         string
	DynamoTableName string
	DynamoRegion    string
	GrpcAddress     string
}

func main() {

	var s Spec
	err := envconfig.Process("service", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	service.Run(s.Address, s.DynamoTableName, s.DynamoRegion, s.GrpcAddress)
}
