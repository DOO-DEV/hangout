ENTRY_POINT?=./main.go

setup:
	@echo Installing dependencies...
	go mod tidy
	go install github.com/cosmtrek/air@latest
	go install github.com/swaggo/swag/cmd/swag@latest # if not working change the version

dev:
	air -c .air.toml

migration-up:
	@echo migrating up...
	go run cmd/migration/main.go "migrate-up"

migration-down:
	@echo migrating down...
	go run cmd/migration/main.go "migrate-down"

generate-doc:
	swag fmt && swag init -g ${ENTRY_POINT} -o ./docs --generatedTime=true