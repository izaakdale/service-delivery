PROTO_DIR=model/delivery

server:
	go build -o bin/blog/server ./blog/server
client:
	go build -o bin/blog/client ./blog/client

gproto:
	protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ${PROTO_DIR}/*.proto \
	 --go-grpc_out=. --go-grpc_opt=paths=source_relative ${PROTO_DIR}/*.proto