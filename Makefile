postgres_url?="postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable"

# start postgres as docker container
.PHONY: docker
docker:
	docker build -t softarch-postgres-image .
	docker run -d --name softarch-postgres -p 5432:5432 softarch-postgres-image

# sqlc for generating queries and schemas in Golang from raw SQL
.PHONY: sqlc
sqlc:
	docker run --rm -v D:\workspaces\SOFTARCH\SOFT-ARCH-REST-HW-6331313021:/src -w /src kjconroy/sqlc generate

### migrate schemas

# createnew migration file
.PHONY: migrate-new
migrate-new:
	migrate create -ext sql -dir database/migrations -seq ${name}

# migrate postgres up
.PHONY: migrate-up
migrate-up:
	migrate -database ${postgres_url} -path database/migrations up

# migrate postgres down
.PHONY: migrate-down
migrate-down:
	migrate -database ${postgres_url} -path database/migrations down

# migrate fix version
.PHONY: migrate-fix
migrate-fix:
	migrate -database ${postgres_url} -path database/migrations force ${version}\

# migrate drop everything, VERY DANGEROUS
.PHONY: migrate-reset
migrate-reset:
	migrate -database ${postgres_url} -path database/migrations drop