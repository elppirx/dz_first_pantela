DB_DSN := "postgres://postgres:test123@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	${MIGRATE} up

migrate-down:
	${MIGRATE} down

run:
	go run cmd/app/main.go

gen:
	oapi-codegen -config openapi/.openapi -include-tags messages -package messages openapi/openapi.yaml > ./iternal/web/messages/api.gen.go

gen-users:
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./iternal/web/users/api.gen.go

lint:
	golangci-lint run --out-format=colored-line-number