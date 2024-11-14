all: test build

build:
	cd file-encryption; \
	docker compose up --build

test:
	echo "running file encryption tests..."; \
	cd file-encryption; \
	go test ./... || exit 1