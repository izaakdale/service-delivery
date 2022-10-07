package dao

import "github.com/izaakdale/service-delivery/model/delivery"

type DeliveryRecord struct {
	OrderId    string `dynamodbav:"PK"`
	RecordType string `dynamodbav:"SK"`
	Status     string
	Address    *delivery.Address
}
