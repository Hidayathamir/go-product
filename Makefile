# Lint using golangci-lint.
lint:
	golangci-lint run --config .golangci.yml ./...

###################################

# Remove docker image with tag None.
clear-none-docker-images:
	docker images --filter "dangling=true" -q --no-trunc | xargs docker rmi

###################################

# Run postgres and redis container.
compose-up-postgres-redis:
	docker compose up -d go-profile-db-postgres go-profile-redis

compose-down-postgres-redis:
	docker compose down go-profile-db-postgres go-profile-redis

# Run go app.
go-run:
	go run main.go -include-migrate

# or run go using air (live reload golang).
air:
	air -c .air.toml

# Run test unit.
go-test-unit:
	go clean -testcache && \
	go test -cover ./internal/... -run TestUnit

###################################

# For deployment. Run postgres container also build and run go app container.
deploy:
	docker compose up --build

undeploy:
	docker compose down

###################################

# Generate proto file.
generate-proto:
	protoc \
		--go_out=.      --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pkg/goproductgrpc/*.proto
