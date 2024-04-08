postgres:
	docker run --name postgres_bank --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres

createdb:
	docker exec -it postgres_bank createdb bank -U root -O root

dropdb:
	docker exec -it postgres_bank dropdb bank -U root

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	./main

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/narymbaev/simple-bank/db/sqlc Store

build:
	go build main.go

.PHONY: createdb dropdb postgres migrateup migratedown migratedown1 sqlc test server mock build