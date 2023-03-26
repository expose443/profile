postgres:
	 docker stop postgres
	 docker run -d --rm --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=password postgres:latest

compose:
	docker-compose up --build

down:
	docker-compose down