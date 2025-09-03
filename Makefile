postgresinit:
	docker-compose up -d

postgresdown:
	docker-compose down

postgres:
	docker exec -it postgres psql -U postgres -t go_jwt_oauth2
	
.PHONY: postgresinit postgres postgresdown
