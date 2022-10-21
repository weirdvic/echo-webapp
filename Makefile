.PHONY: run
run:
	go run ./cmd/web

.PHONY: new-migration
new-migration:
	migrate create -ext sql -dir migrations/ -seq -digits 6 $(name)

.PHONY: migrate-up
migrate-up:
	migrate -database ${DB_DSN} -path migrations/ up

.PHONY: migrate-down
migrate-down:
	migrate -database ${DB_DSN} -path migrations/ down
