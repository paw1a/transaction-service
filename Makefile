build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

build-macos:
	go mod download && go build -o ./.bin/app ./cmd/app/main.go

run: build
	docker-compose up

clean:
	rm -rf .bin .data

.DEFAULT_GOAL := run
