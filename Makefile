BINARY_NAME=rasantApp

build:
	@go mod vendor
	@echo "Building Rasant..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Rasant built!"

run: build
	@echo "Starting Rasant..."
	@./tmp/${BINARY_NAME} &
	@echo "Rasant started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

start_compose:
	docker-compose up -d

stop_compose:
	docker-compose down

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Rasant..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Rasant!"

restart: stop start