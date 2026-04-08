CC=go
BINARY_NAME=vex

all: build

build:
	$(CC) build -o $(BINARY_NAME) main.go

test:
	$(CC) test ./...

clean:
	rm -f $(BINARY_NAME)
	rm -f report.json

docker-build:
	docker build -t $(BINARY_NAME) .

run-challenge:
	cd ../vex-challenge && go run main.go

help:
	@echo "Vex Build System"
	@echo "  make build          - Build the binary"
	@echo "  make clean          - Remove binary and reports"
	@echo "  make docker-build   - Build Docker image"
	@echo "  make run-challenge  - Start the local BOLA challenge lab"
