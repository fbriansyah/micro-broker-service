FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o broker-microservice ./cmd/

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/broker-microservice .
COPY app.env .

EXPOSE 8081

ENTRYPOINT [ "/app/broker-microservice" ]