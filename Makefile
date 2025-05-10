generateswag:
	cd server && swag init -g cmd/api-gateway/main.go

compose_rebuild:
	docker compose up --build -d

danek:
	cd server && swag init -g cmd/api-gateway/main.go && cd .. && docker compose up --build -d