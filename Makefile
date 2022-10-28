PROTO_DIR=model/delivery

run:
	SERVICE_ADDRESS=8081 SERVICE_DYNAMOTABLENAME=ordering-app SERVICE_DYNAMOREGION=eu-west-2 SERVICE_GRPCADDRESS=localhost:50001 go run main.go

build:
	GOOS=linux go build -o service-delivery . 
docker_build:
	docker build --no-cache -t izaakdale/service-delivery .
build_amd:
	GOOS=linux GOARCH=amd64 go build -o service-delivery . 
docker_build_amd:
	docker buildx build --no-cache --platform linux/amd64  -t izaakdale/service-delivery .

gproto:
	protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ${PROTO_DIR}/*.proto \
	 --go-grpc_out=. --go-grpc_opt=paths=source_relative ${PROTO_DIR}/*.proto
