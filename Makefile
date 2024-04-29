.PHONY: build test run

build:
	@echo "Building the application..."
	@mkdir -p bin
	@go build -o bin/TicketIt cmd/main.go

test:
	@echo "Running tests..."
	@go test -v ./...

run: build
	@echo "Running the application..."
	@./bin/TicketIt

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down