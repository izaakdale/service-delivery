package dao

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func StoreAddress(a *AddressRecord) error {
	statusMap, err := dynamodbattribute.MarshalMap(a)
	if err != nil {
		return err
	}

	writeRequest := []*dynamodb.WriteRequest{}
	writeRequest = append(writeRequest, &dynamodb.WriteRequest{
		PutRequest: &dynamodb.PutRequest{
			Item: statusMap,
		},
	})

	_, err = client.BatchWriteItemWithContext(context.Background(), &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			client.tableName: writeRequest,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
