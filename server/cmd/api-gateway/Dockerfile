FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api-gateway


FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app .

ARG PORT=10000
ENV PORT=${API_GATEWAY_PORT}
EXPOSE $PORT

CMD ["./api"]
