run:
	go run ./cmd/main.go

down:
	docker-compose down

up:
	docker-compose up -d

build-image:
	docker build -t efner/broker-microservice:1.0 .

.PHONY: run down up