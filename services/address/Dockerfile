# Build stage
FROM golang:1.23.4-alpine AS builder

WORKDIR /src
COPY . .

# COPY go.mod go.sum ./
RUN go build -C cmd/api -o /out/address-svc

# Final stage
FROM alpine:3.19

WORKDIR /app
COPY --from=builder /out/address-svc /app/address-svc

EXPOSE 8080
ENTRYPOINT ["/app/address-svc"]