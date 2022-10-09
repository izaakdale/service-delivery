package service

import (
	"github.com/izaakdale/service-delivery/dao"
	"github.com/izaakdale/utils-go/logger"
	"github.com/julienschmidt/httprouter"
)

type Service struct {
	Router *httprouter.Router
}

func NewService() *Service {
	logger.Info("Starting service...🚀")

	// TODO remove hard coding
	dao.Init("ordering-app", "eu-west-2")

	return &Service{
		Router: Routes(),
	}
}

func Routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc("POST", "/status", UpdateDeliveryStatusHandler)
	return router
}
