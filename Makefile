
deps:
	go mod tidy
	go mod vendor

lint:
	gosec ./...
	go vet ./...