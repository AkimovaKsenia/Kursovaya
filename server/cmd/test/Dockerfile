FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o test ./cmd/test


FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app .

ARG PORT=10001
ENV PORT=${FILM_SERVICE_PORT}
EXPOSE $PORT

CMD ["./test"]
