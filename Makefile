MAKEFLAGS += --silent

all: test build run

# ADD YOUR MICROSERVICE FOLDER HERE
SERVICES=file-encryption file-transfer accounts

build:
	echo "Building services..."; \
	docker compose build

run:
	echo "Running services..."; \
	docker compose up -d

stop:
	echo "Stopping services..."; \
	docker compose down

install:
	cd accounts; \
	go mod download; \
	cd ../file-transfer; \
	go mod download; \
	cd ../file-encryption; \
	go mod download; \
	cd ../frontend; \
	npm install --saved

test:
	failed=0; \
	for service in $(SERVICES); do \
		echo "Running $$service tests..."; \
		cd $$service; \
		go test ./... || failed=$$(($$failed + 1)); \
		cd ..; \
	done; \
	echo "Running frontend tests..."; \
	cd frontend; \
	# npm run test:unit || failed=$$(($$failed + 1)); \
	# cd ..; \
	echo "$$failed tests failed."; \
	exit $$failed

test-coverage:
	for service in $(SERVICES); do \
		echo "Running $$service coverage tests..."; \
		cd $$service; \
		go test -cover ./...; \
		cd ..; \
	echo "Running frontend coverage tests..."; \
	done; \
	cd frontend; \
	# npm run test:coverage; \
