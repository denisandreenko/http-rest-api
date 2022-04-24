.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: migrate-up
migrate-up:
	migrate -path migrations -database "postgres://localhost/restapi_dev?sslmode=disable" up

.PHONY: migrate-test-up
migrate-test-up:
	migrate -path migrations -database "postgres://localhost/restapi_test?sslmode=disable" up

.PHONY: migrate-up-docker
migrate-up-docker:
	migrate -path migrations -database "postgres://postgres:root@localhost/restapi_dev?sslmode=disable" up

.PHONY: migrate-down
migrate-down:
	migrate -path migrations -database "postgres://localhost/restapi_dev?sslmode=disable" down

.PHONY: migrate-test-down
migrate-test-down:
	migrate -path migrations -database "postgres://localhost/restapi_test?sslmode=disable" down

.PHONY: migrate-down-docker
migrate-down-docker:
	migrate -path migrations -database "postgres://postgres:root@localhost/restapi_dev?sslmode=disable" down

.DEFAULT_GOAL := build
