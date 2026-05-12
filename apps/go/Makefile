APP_NAME      := server
BUILD_DIR     := bin
MAIN_PATH     := cmd/server/main.go
DOCKER_IMAGE  := go
DOCKER_TAG    := latest

.PHONY: help run build clean test coverage lint fmt tidy docker-build

run:
	@env $$(grep -v '^#' .env | xargs) go run $(MAIN_PATH)

clean:
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out

tidy:
	@go mod tidy
	@go mod verify

fmt:
	@go fmt ./...

lint:
	@golangci-lint run ./...

test:
	@go test -v -race -short ./...

coverage:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

build: clean
	@CGO_ENABLED=0 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)

docker-build:
	@docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
