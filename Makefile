default: run

run:
	go run cmd/main.go

build:
	go build cmd/main.go

docker:
	docker/sudo docker-compose up -d