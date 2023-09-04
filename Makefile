init:
	docker run --name examples-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=examples -d postgres:12-alpine

postgresStart:
	docker start hackhack-postgres

postgresStop:
	docker stop hackhack-postgres

migrateup:
	migrate -path cmd/migrations -database "postgresql://root:postgres@localhost:5432/examples?sslmode=disable" -verbose up

migrateup1:
	migrate -path cmd/migrations -database "postgresql://root:postgres@localhost:5432/examples?sslmode=disable" -verbose up 1

migratedown:
	migrate -path cmd/migrations -database "postgresql://root:postgres@localhost:5432/examples?sslmode=disable" -verbose down

migratedown1:
	migrate -path cmd/migrations -database "postgresql://root:postgres@localhost:5432/examples?sslmode=disable" -verbose down 1

run:
	go run ./cmd/app/main.go

.PHONY: init migrateup migrateup1 migratedown migratedown1 run postgresStart postgresStop