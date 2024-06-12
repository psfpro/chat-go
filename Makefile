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

gen-keys:
	openssl genrsa -out ./config/jwt/private.pem 2048
	openssl rsa -pubout -in ./config/jwt/private.pem -out ./config/jwt/public.pem