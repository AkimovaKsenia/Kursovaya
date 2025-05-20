generateswag:
	cd server && swag init -g cmd/api-gateway/main.go

container_build:
	docker build -t gewinn2/cinema_mastery_api-gateway:latest -f ./server/cmd/api-gateway/Dockerfile ./server
	docker build -t gewinn2/cinema_mastery_film_service:latest -f ./server/cmd/film_service/Dockerfile ./server
	docker build -t gewinn2/cinema_mastery_cinema_service:latest -f ./server/cmd/cinema_service/Dockerfile ./server
	docker build -t gewinn2/cinema_mastery_frontend:latest ./client

container_push:
	docker image push gewinn2/cinema_mastery_api-gateway:latest
	docker image push gewinn2/cinema_mastery_film_service:latest
	docker image push gewinn2/cinema_mastery_cinema_service:latest
	docker image push gewinn2/cinema_mastery_frontend:latest

danek:
	cd server && swag init -g cmd/api-gateway/main.go && cd .. && docker compose up --build -d