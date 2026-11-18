.SILENT:
run:
	go run ./cmd/main.go

test:
	go test -v -race ./...

migrate_up:
	migrate -path ./schema -database "postgres://anme:1234@localhost:5432/postgres?sslmode=disable" up

migrate_down:
	migrate -path ./schema -database "postgres://anme:1234@localhost:5432/postgres?sslmode=disable" down