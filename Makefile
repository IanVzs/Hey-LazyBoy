# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GoTool=$(GOCMD) tool
BINARY_NAME=snowflakes
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build:
	@echo "start go build...."
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
lint:
	golint ./...
tool:
	$(GoTool) vet . |& grep -v vendor; true
swag:
	swag init
install_swag:
	go install github.com/swaggo/swag/cmd/swag@latest
clean:
	$(GOCLEAN) -i .
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
swagger:
	swag init
deps:
	$(GOGET) github.com/go-redis/redis/v8
	$(GOGET) github.com/natefinch/lumberjack
	$(GOGET) go.uber.org/zap
help:
	@echo "make: test & compile packages and dependencies"
	@echo "make run: make & run"
	@echo "make tool: run specified go tool"
	@echo "make lint: ..."
	@echo "make clean: remove object files and cache files"
	@echo "make help: show this"
