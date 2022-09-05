createctn:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:12.3-alpine

killctn:
	docker kill postgres12

rmctn:
	docker rm postgres12 -f

startctn:
	docker start postgres12