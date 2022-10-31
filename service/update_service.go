package service

import (
	"github.com/izaakdale/utils-go/logger"
	"github.com/julienschmidt/httprouter"
)

type Service struct {
	Router *httprouter.Router
}

func NewService() *Service {
	logger.Info("Starting service...🚀")

	return &Service{
		Router: Routes(),
	}
}

func Routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc("POST", "/status", UpdateDeliveryStatusHandler)
	return router
}
