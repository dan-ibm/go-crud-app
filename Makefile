db.start:
	sudo docker-compose up -d

db.stop:
	sudo docker-compose stop

db.bash:
	sudo docker-compose exec db bash