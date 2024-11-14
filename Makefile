all: test build

# ADD YOUR MICROSERVICE FOLDER HERE
SERVICES=file-encryption accounts

build:
	echo "Building services..."; \
	docker compose up --build

test:
	failed=0; \
	for service in $(SERVICES) ; do \
		echo "Running $$service tests..."; \
		cd $$service; \
		go test ./... || failed=$$(($$failed + 1)); \
		cd ..; \
	done; \
	echo "$$failed tests failed."; \
	exit $$failed
