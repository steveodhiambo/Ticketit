build:
	@go build -o bin/TicketIt cmd/main.go

test:
	@go test -v ./..

run:
	@./bin/TicketIt