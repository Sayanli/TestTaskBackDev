.PHONY: run
run:
	go run cmd/app/main.go

.PHONY: compose-up
compose-up:
	docker-compose up -d