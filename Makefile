MYSQL_USERNAME=root
MYSQL_PASSWORD=pw
MYSQL_HOST=localhost
MYSQL_PORT=3306

# Run all up migrations following the most recently run migration
db-up:
	migrate -path db/migrations -database 'mysql://${MYSQL_USERNAME}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/db' up

# Run all down migrations, effectively dropping all tables
db-down:
	migrate -path db/migrations -database 'mysql://${MYSQL_USERNAME}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/db' down

# Force-drop all tables
db-drop:
	migrate -path db/migrations -database 'mysql://${MYSQL_USERNAME}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/db' drop

# Run the application
run:
	go run cmd/api/main.go

# Run the tests
unit-test:
	go test ./...

# Insert some data for testing into the db
seed:
	go run cmd/seed/main.go

# Regenerate mocks for use in unit testing
.PHONY: mocks
mocks:
	mockery --all
