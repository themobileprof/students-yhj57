# Variables
APP_NAME = student-management-api
GO = go
GO_BUILD = $(GO) build
GO_RUN = $(GO) run
GO_CLEAN = $(GO) clean
GO_TEST = $(GO) test

# Default target
all: build

# Build the application
build:
	$(GO_BUILD) -o $(APP_NAME) main.go

# Run the application
run:
	$(GO_RUN) main.go

# Clean the build
clean:
	$(GO_CLEAN)
	rm -f $(APP_NAME)

# Run tests
test:
	$(GO_TEST) ./...

.PHONY: all build run clean test