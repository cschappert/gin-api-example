MYSQL_USERNAME=root
MYSQL_PASSWORD=pw
MYSQL_HOST=localhost
MYSQL_PORT=3306

db-up:
	migrate -path db/migrations -database 'mysql://${MYSQL_USERNAME}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/db' up

db-down:
	migrate -path db/migrations -database 'mysql://${MYSQL_USERNAME}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/db' down

db-drop:
	migrate -path db/migrations -database 'mysql://${MYSQL_USERNAME}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/db' drop

run:
	go run cmd/api/main.go
