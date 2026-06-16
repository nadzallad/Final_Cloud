FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum* ./

RUN go mod download 2>/dev/null || true

COPY . .

RUN go mod tidy && go build -o shipment-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/shipment-service .

EXPOSE 8085

CMD ["./shipment-service"]