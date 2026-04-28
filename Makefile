GOOSE_DRIVER=mysql
GOOSE_DBSTRING=root:root@tcp(localhost:3306)/ecommerce
GOOSE_MIGRATION_DIR=./sql/schema

# name app
APP_NAME := server

# run app
dev:
	go run ./cmd/${APP_NAME}
run:
	docker-compose up -d && go run./cmd/${APP_NAME}
kill:
	docker-compose kill
up:
	docker-compose up -d
down:
	docker-compose down
upse:
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up
downse:
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down
resetse:
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) reset

.PHONY: run downse upse resetse

.PHONY: air