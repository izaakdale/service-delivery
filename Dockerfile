FROM golang:1.18.3-alpine3.16 AS builder
WORKDIR /build
COPY . .
RUN GOOS=linux go build -o service-delivery . 

# move binary file to lightweight image
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /build/service-delivery .

EXPOSE 80

CMD ["./service-delivery"]