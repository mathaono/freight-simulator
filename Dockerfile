# Build stage
FROM golang:1.23.4-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# COPY go.mod go.sum ./
RUN go build -o /out/address-svc ./services/address/cmd/api

# Final stage
FROM alpine:3.19

WORKDIR /app
COPY --from=builder /out/address-svc /app/address-svc

EXPOSE 8080
ENTRYPOINT ["/app/address-svc"]