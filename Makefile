up:
	docker compose up -d

down:
	docker compose down

check: clear vet test

clear:
	clear

vet:
	go vet ./...

test:
	go test -v ./...

protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
	  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	  proto/chatgo.proto