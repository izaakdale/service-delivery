package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/izaakdale/service-delivery/dao"
)

func Run(addr, tableName, region, grpcSrvAddr string) {
	go RunGrpcServer(grpcSrvAddr)

	dao.Init(tableName, region)
	service := NewService()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", addr), service.Router))
}
