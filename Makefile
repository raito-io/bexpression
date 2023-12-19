# To try different version of Go
GO := go

gotestsum := go run gotest.tools/gotestsum@latest

generate:
	go generate ./...

build: generate
	go build ./...

test:
	$(gotestsum) --debug --format testname -- -race -mod=readonly -v ./...

test-coverage:
	$(gotestsum) --debug --format testname -- -race -mod=readonly -v -coverpkg=./... -covermode=atomic -coverprofile=coverage.txt ./...
	go tool cover -html=unittest-coverage.txt -o integrationtest-coverage.html

lint:
	golangci-lint run ./...
	go fmt ./...