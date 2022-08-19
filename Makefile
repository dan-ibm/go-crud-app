init:
	cp .env.example .env
	make db.start
	make run

db.start:
	sudo docker-compose up -d

db.stop:
	sudo docker-compose stop

db.bash:
	sudo docker-compose exec db bash

run:
	go run cmd/app/main.go

swag:
	swag init -g cmd/app/main.go