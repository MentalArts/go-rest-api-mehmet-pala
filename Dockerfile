# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.2
RUN swag init -g cmd/main.go
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs
EXPOSE 8080
CMD ["./main"]