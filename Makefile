createdb: 
	docker exec -it postgres12 createdb --username=root --owner=root owe

dropdb: 
	docker exec -it postgres12 dropdb owe

postgres: 
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=harsh -d postgres:12-alpine

migrateup: 
	migrate -path db/migration -database "postgres://root:harsh@localhost:5432/owe?sslmode=disable" -verbose up

migratedown: 
	migrate -path db/migration -database "postgres://root:harsh@localhost:5432/owe?sslmode=disable" -verbose down

sqlcgen:
	sqlc generate

test:
	go test -v -cover ./...

# change this manually todo: make this dynamic
createmigration: 
	migrate create -ext .sql -dir ./db/migration add-balance-field