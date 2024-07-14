# Project variables
PROJECT_NAME := speech_to_text
PKG := ./cmd/...
MAIN := ./cmd/main.go

# Go commands
BUILD := go build
CLEAN := go clean
FMT := go fmt
VET := go vet
TEST := go test
RUN := go run

# Targets
.PHONY: all build clean fmt vet test run

all: fmt vet test build

build:
	$(BUILD) -o $(PROJECT_NAME) $(MAIN)

clean:
	$(CLEAN)
	rm -f $(PROJECT_NAME)

fmt:
	$(FMT) $(PKG)

vet:
	$(VET) $(PKG)

test:
	$(TEST) $(PKG)

run:
	$(RUN) $(MAIN)

