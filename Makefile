all: test build

# ADD YOUR MICROSERVICE FOLDER HERE
SERVICES=file-encryption

build:
	for service in $(SERVICES) ; do \
		echo "Building $$service..."; \
  		cd $$service; \
		docker compose up --build; \
	done

test:
	failed=0; \
	for service in $(SERVICES) ; do \
		echo "Running $$service tests..."; \
		cd $$service; \
		go test ./... || failed=$$(($$failed + 1)); \
		cd ..; \
	done; \
	echo "$$failed tests failed."; \
	exit $$failed;
