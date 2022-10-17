install:
	docker run -d -p 5432:5432 --name postgres --restart=always -e POSTGRES_USER=aditya.kristianto@mncgroup.com -e POSTGRES_PASSWORD=hacktiv8 -d postgres
	docker run -d -p 8081:80 --name=pgadmin --restart=always -e PGADMIN_DEFAULT_EMAIL=aditya.kristianto@mncgroup.com -e PGADMIN_DEFAULT_PASSWORD=hacktiv8 dpage/pgadmin4 

run:
	go run cmd/final-project/main.go

swag-init:
	swag init --parseInternal -g cmd/final-project/main.go