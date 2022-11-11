build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

run:
	docker-compose up db app

clean:
	rm -rf .bin .data

.DEFAULT_GOAL := run
