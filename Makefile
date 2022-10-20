
deps:
	go mod tidy
	go mod vendor

lint:
	@staticcheck ./...