PROJECTNAME = credit-scoring-service

lint:
	golangci-lint run  --config=.golangci.yml --timeout=180s ./...

generate:
	go generate ./..

run-migrate-local:
	sql-migrate up -env="development" -limit=0

stop-migrate-local:
	sql-migrate down -env="development" -limit=0

build:
	go build -o ./build/${PROJECTNAME} ./cmd/${PROJECTNAME}/main.go || exit 1

run:
	go run ./cmd/${PROJECTNAME}/main.go

# enter-db:
# 	docker exec -it sport-crm-api-as-db-1 psql db db

# docker-up:
# 	docker compose up -d

# docker-stop:
# 	docker compose stop