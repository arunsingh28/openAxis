# variabels

APP_NAME  = api
CMD_NAME  = $(APP_NAME)
PKG_NAME  = github.com/$(ORG_NAME)/$(APP_NAME)
PKG_PATH  = $(GOPATH)/src/$(PKG_NAME)
MAIN_FILE = $(PKG_PATH)/cmd/$(CMD_NAME)/main.go

.PHONY: all build run clean test

# default target
all: build

# build the app
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(CMD_NAME) $(MAIN_FILE)

# run the app
run:
	@echo "Running $(APP_NAME)..."
	@go run $(MAIN_FILE)

# clean the app
clean:
	@echo "Cleaning $(APP_NAME)..."
	@rm -f $(CMD_NAME)

# test the app
test:
	@echo "Testing $(APP_NAME)..."
	@go test -v ./...

