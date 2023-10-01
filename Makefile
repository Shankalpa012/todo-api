.PHONY: migrate

run:
	nodemon --exec "go run" main.go

migrate:
	go run migrate/migration.go