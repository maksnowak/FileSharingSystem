all: test build

# ADD YOUR MICROSERVICE FOLDER HERE
SERVICES=file-encryption file-transfer accounts

build:
	echo "Building services..."; \
	docker compose up --build -d

stop:
	echo "Stopping services..."; \
	docker compose down

test:
	failed=0; \
	for service in $(SERVICES); do \
		echo "Running $$service tests..."; \
		cd $$service; \
		go test ./... || failed=$$(($$failed + 1)); \
		cd ..; \
	done; \
	echo "$$failed tests failed."; \
	exit $$failed

test-coverage:
	for service in $(SERVICES); do \
		echo "Running $$service coverage tests..."; \
		cd $$service; \
		go test -cover ./...; \
		cd ..; \
	done; \
