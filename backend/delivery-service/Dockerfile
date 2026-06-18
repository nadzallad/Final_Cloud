FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum* ./

RUN go mod download 2>/dev/null || true

COPY . .

RUN go mod tidy && go build -o delivery-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/delivery-service .

EXPOSE 8086

CMD ["./delivery-service"]