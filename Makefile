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
