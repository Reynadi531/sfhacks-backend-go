build:
	go build -o bin/main main.go

run:
	go run main.go

docker-dev:
	docker-compose -f docker/docker-compose.dev.yaml up

docker-dev-build:
	docker-compose -f docker/docker-compose.dev.yaml up --build

docker-dev-standalone:
	docker-compose -f docker/docker-compose.dev-standalone.yaml up

docker-build:
	docker build . -f docker/Dockerfile -t sfhacks-go:latest

dev:
	air
