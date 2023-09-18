dev:
	air -c .air.toml

migration-up:
	@echo migrating up...
	go run cmd/migration/main.go "migrate-up"

migration-down:
	@echo migrating down...
	go run cmd/migration/main.go "migrate-down"