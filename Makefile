# Variables
DESTDIR ?= /usr/local
BINDIR ?= $(DESTDIR)/bin
GOPATH ?= $(shell go env GOPATH)

# Default target
all: build

build:
	@echo "Building gnosis..."
	go build -o gnosis main.go utils.go

clean:
	@echo "Cleaning up..."
	go clean

install: build
	@echo "Installing gnosis to $(BINDIR)..."
	mkdir -p $(BINDIR)
	cp -f gnosis $(BINDIR)
	chmod 755 $(BINDIR)/gnosis

uninstall:
	@echo "Uninstalling gnosis from $(BINDIR)..."
	rm -f $(BINDIR)/gnosis

test:
	@echo "Running tests..."
	go test

help:
	@echo "Available targets:"
	@echo "  all (default)    - Build the project"
	@echo "  build            - Build the project"
	@echo "  clean            - Remove built binaries"
	@echo "  install          - Install the project"
	@echo "  uninstall        - Uninstall the project"
	@echo "  test             - Run tests"
	@echo "  help             - Display this help message"
