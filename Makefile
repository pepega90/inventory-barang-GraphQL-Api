run:
	go run cmd/*.go

up:
	sudo docker-compose up

down:
	sudo docker-compose down

generate:
	go generate ./...