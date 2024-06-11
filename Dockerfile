# Build the application from source
FROM golang:1.22.2 AS build-stage
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY .env .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/TicketIt cmd/main.go

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM scratch AS build-release-stage
WORKDIR /

COPY --from=build-stage /app/bin/TicketIt .
COPY --from=build-stage /app/.env .

EXPOSE 4000

ENTRYPOINT ["./TicketIt"]