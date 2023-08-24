.PHONY: run
run:
	go run cmd/app/main.go

.PHONY: compose-up
compose-up:
	docker-compose up -d

.PHONY: swag
swag:
	swag init -g internal/app/app.go --parseInternal --parseDependency