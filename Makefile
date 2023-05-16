API_PATH=api/notification
PROTO_OUT_DIR=pkg/notificationapi
PROTO_API_DIR=$(API_PATH)
ARGS=-fix

.PHONY: gen
gen: gen-proto generate

.PHONY: gen-proto
gen-proto:
	mkdir -p $(PROTO_OUT_DIR)
	protoc \
		-I $(API_PATH) \
    	--go_out=pkg/notificationapi --go_opt=paths=source_relative \
    	--go-grpc_out=pkg/notificationapi  --go-grpc_opt=paths=source_relative \
    ./$(PROTO_API_DIR)/v1/*.proto

evans:
	evans --port 9001 -r repl

test:
	go test ./... -cover -count=1

generate:
	go generate ./...

lint:
	go/lint proto/lint

go/lint:
	golangci-lint run  --config=.golangci.yml --timeout=30s ./...

proto/lint:
	protolint lint $(ARGS) $(PROTO_API_DIR)/*

run:
	go run cmd/notification/main.go

r-up:
	docker run --name rabbit -e RABBITMQ_DEFAULT_USER=rabbitmq -e RABBITMQ_DEFAULT_PASS=rabbitmq -p 15672:15672 -p 5672:5672 -d --rm rabbitmq:3-management

r-stop:
	docker stop rabbit
