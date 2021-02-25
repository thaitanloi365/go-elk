start:
	docker-compose --env-file=./.env up --build $(args)

startd:
	docker-compose --env-file=./.env up --build -d $(args)