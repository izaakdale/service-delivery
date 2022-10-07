package dao

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var client *Client

type DynamoDb interface {
	dynamodbiface.DynamoDBAPI
}

type Client struct {
	dynamodbiface.DynamoDBAPI
	tableName string
}

func Init(tableName, region string) {

	session, err := session.NewSession(aws.NewConfig().WithRegion(region))
	if err != nil {
		log.Fatalf("Failed to created a session for the db client... Error: %v", err)
	}

	dynamo := dynamodb.New(session)

	client = &Client{
		DynamoDBAPI: dynamo,
		tableName:   tableName,
	}
}
