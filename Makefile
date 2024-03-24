# Define variables
GOBIN := ./bin
GOPACKAGES := $(shell go list ./... | grep -v /vendor)

# Build command
build:
	go build -o $(GOBIN)/navi ./main.go

# Run command
run: build
	$(GOBIN)/navi

# Test command (assuming you have unit tests)
test:
	go test -v $(GOPACKAGES)

# Clean command
clean:
	rm -rf $(GOBIN)

# Live reload command (optional)
livereload:
	# Replace with your live reload command (e.g., go-watcher)
	# go-watcher -n -o $(GOBIN)/navi ./...

gomod:
	go mod tidy

vendor:
	go mod vendor

# Help command
help:
	@echo "Available commands:"
	@echo "  build     - Build the Navi executable"
	@echo "  run       - Run the Navi program"
	@echo "  test      - Run unit tests"
	@echo "  clean     - Remove built files"
	@echo "  livereload - Run with live reloading (optional)"
	@echo "  help      - Display this help message"
	@echo "  gomod     - Update go.mod file"
	@echo "  vendor    - Update vendor directory"
