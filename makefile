run:
	@echo "Running the server..."
	@go run ./cmd/server

build:
	@echo "Building the binary..."
	@go build -o bin/server ./cmd/server

test:
	@echo "Running tests..."
	
clean:
	@echo "Cleaning up..."
	@rm -f bin/server