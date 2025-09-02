.PHONY: tidy build run

tidy:
	cd services/address && go mod tidy

build:
	cd services/address/cmd/api && go build -o ../../bin/address-svc

run:
	cd services/address/cmd/api && go run .

test:
	cd services/address && go test ./...