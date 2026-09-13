.PHONY: run test fmt build compose-up compose-down
run:
	go run ./cmd/server
test:
	go test ./...
fmt:
	gofmt -w $$(find . -name '*.go' -type f)
build:
	go build -o bin/cloud-platform-api ./cmd/server
compose-up:
	docker compose up --build -d
compose-down:
	docker compose down
