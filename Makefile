build:
	@go build -o bin/ticketit cmd/main.go

test:
	@go test -v ./..

run:
	@./bin/ticketit