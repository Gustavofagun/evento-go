FROM golang:latest as app
WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]


FROM redis:7.0.11-alpine3.17 as redis
EXPOSE 6379
CMD ["redis-server"]
 