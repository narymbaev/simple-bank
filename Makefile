postgres:
	docker run --name postgres_bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres

createdb:
	docker exec -it postgres_bank createdb bank -U root -O root

dropdb:
	docker exec -it postgres_bank dropdb bank -U root

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test