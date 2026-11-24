start-db:
	docker compose -f docker-compose.dev.yaml up -d

stop-db:
	docker compose -f docker-compose.dev.yaml down

start:
	SERVER_ADDRESS=":8100" \
	DATABASE_DSN="host=localhost port=5500 user=postgres password=postgres dbname=tunelport-svc-devices_test sslmode=disable" \
	go run cmd/grpc/main.go

test:
	go clean -testcache && go test  ./...


app-proto:
	bash bin/app/proto.sh