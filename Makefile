.PHONY: build test lint migrate

build:
	go build -v ./...

test:
	go test -v ./...

watch:
	watchexec --exts go "gotestsum -- -p 1 ./..."

lint:
	golint ./...

migrate:
	go run ./migrations/main.go

run:
	go run ./cmd/user_interaction/main.go

init_session:
	go run ./cmd/initial_session/main.go