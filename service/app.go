package service

import (
	"log"
	"net/http"
)

func Run() {
	go RunGrpcServer()

	service := NewService()
	log.Fatal(http.ListenAndServe("localhost:8081", service.Router))
}
