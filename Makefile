run:
	go run ./cmd/main.go

down:
	docker-compose down

up:
	docker-compose up -d

.PHONY: run down up