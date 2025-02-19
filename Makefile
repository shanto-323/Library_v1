run: clean build
	@./bin/app

build:
	@go build -o bin/app ./cmd

clean:
	@rm -rf ./bin/app

GOOSE := goose
MIGRATIONS_DIR := ./migrations
DB_DRIVER := postgres
DSN := "user=postgres password=1234 dbname=library host=localhost port=5432 sslmode=disable"

status:
	$(GOOSE) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DSN) status

up:
	$(GOOSE) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DSN) up

down:
	$(GOOSE) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DSN) down

reset: down up
