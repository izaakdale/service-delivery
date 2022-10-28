# build a tiny docker image
FROM alpine:latest 

RUN mkdir /app
# requires a built binary
COPY service-delivery /app

CMD ["/app/service-delivery"]