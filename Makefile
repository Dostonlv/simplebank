postgres:
	docker run --name Dostonlv -p 5432:5432 -e POSTGRES_USER=dostonlv -e POSTGRES_PASSWORD=dostonlv -d postgres

createdb:
	docker exec -it Dostonlv createdb --username=dostonlv --owner=dostonlv simple_bank

dropdb:
	docker exec -it Dostonlv dropdb --username=dostonlv simple_bank

migrateup:
	migrate --path db/migration --database "postgresql://dostonlv:secret@localhost:5432/simple_bank?sslmode=disable" --verbose up


migratedown:
	migrate --path db/migration --database "postgresql://dostonlv:secret@localhost:5432/simple_bank?sslmode=disable" --verbose down

sqlc:
	sqlc generate
test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb  -destination db/mock/store.go github.com/Dostonlv/simplebank/db/sqlc Store

.PHONY: cretedb createdb dropdb migrateup migratedown sqlc test server mock
