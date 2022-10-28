package dao

import "github.com/izaakdale/service-delivery/model/delivery"

type AddressRecord struct {
	OrderId    string            `dynamodbav:"PK" json:"order_id,omitempty"`
	RecordType string            `dynamodbav:"SK" json:"record_type,omitempty"`
	Address    *delivery.Address `json:"address,omitempty"`
}

type StatusRecord struct {
	OrderId    string `dynamodbav:"PK" json:"order_id,omitempty"`
	RecordType string `dynamodbav:"SK" json:"record_type,omitempty"`
	Status     string `json:"status,omitempty"`
}
