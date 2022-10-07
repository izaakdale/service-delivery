package dao

import "github.com/izaakdale/service-delivery/model/delivery"

type DeliveryRecord struct {
	OrderId        string `dynamodbav:"PK"`
	DeliveryStatus string `dynamodbav:"SK"`
	Address        *delivery.Address
}
