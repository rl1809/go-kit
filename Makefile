include config/default.env

migrateup:
	@ migrate -path database/mysql/migration -database $(DATABASE_URL) --verbose up

migratedown:
	@ migrate -path database/mysql/migration -database $(DATABASE_URL) --verbose down

sqlc:
	@ sqlc generate

runserver:
	@ go run main.go serve


.PHONY: migrateup migratedown sqlc runserver