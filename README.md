
# TicketIt

TicketIt is a simple ticket management system implemented in Go. The project includes various entities such as users, tickets, comments, settings, statuses, priorities, categories, and audits, which are interconnected as shown in the database schema diagram.

## Database Schema

![Database Schema](path/to/schema/image.png)

## Project Structure

- **cmd/main.go**: The main entry point of the application.
- **cmd/migrate/main.go**: Handles database migrations.
- **cmd/migrate/migrations**: Contains SQL migration files.
- **bin/**: Directory where the compiled application will be placed.

## Makefile Commands

The project includes a Makefile with various commands to build, test, and run the application, as well as to manage database migrations.

### Build the Application

To build the application, run:

```sh
make build
```

This command will compile the Go code and place the executable in the `bin` directory.

### Run the Application

To run the application, use:

```sh
make run
```

This will first build the application and then run the resulting binary.

### Run Tests

To execute the tests, run:

```sh
make test
```

This command will run all the tests in the project.

### Database Migrations

#### Create a New Migration

To create a new migration, use:

```sh
make migration name=your_migration_name
```

This will create a new SQL migration file in the `cmd/migrate/migrations` directory.

#### Apply Migrations

To apply all the pending migrations, run:

```sh
make migrate-up
```

#### Rollback Migrations

To rollback the last applied migration, use:

```sh
make migrate-down
```

## How to Run the Application

1. **Clone the repository:**

```sh
git clone https://github.com/your-username/TicketIt.git
cd TicketIt
```

2. **Build the application:**

```sh
make build
```

3. **Run the database migrations:**

```sh
make migrate-up
```

4. **Run the application:**

```sh
make run
```

## Dependencies

Ensure you have the following dependencies installed:

- Go 1.16+
- `migrate` tool for database migrations

## Contributing

Feel free to fork this repository and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License. See the LICENSE file for details.