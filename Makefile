generateswag:
	cd server && swag init -g cmd/api-gateway/main.go

compose_rebuild:
	docker compose build && docker compose up -d