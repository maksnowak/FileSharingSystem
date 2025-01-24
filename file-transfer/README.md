# File Transfer

## Description

File transfer microservice is responsible for handling CRUD operations of user files.
It communicates with other services to store retrieve specific files.

## Prerequisites

- [Golang version >=1.23](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Configuration

This microservice is dependent on MongoDB database so before starting the web server you need to provide connection configuration to MongoDB database.

### Running MongoDB locally

You can run MongoDB locally with `docker compose` by using `docker compose run mongodb` command in project _file-transfer_ directory.
Then you need to create `.env` file with MongoDB configuration.
Example configuration for MongoDB with user `root` and password `example`:

```bash
MONGODB_URI=mongodb://root:example@localhost:27017/
```

### Connecting to existing instance

Provide connection details in `.env` file.

### Connecting to Storage

#### Azure Storage

You need to specify `AZURE_STORAGE_ACCOUNT_NAME` and `AZURE_STORAGE_ACCOUNT_KEY` in `.env` file.
Example configuration for Azure Storage:

```bash
AZURE_STORAGE_ACCOUNT_NAME=example
AZURE_STORAGE_ACCOUNT_KEY=example
```

#### Local Storage

Alternatively you can use local storage by providing `LOCAL_STORAGE_PATH` and `STORAGE_TYPE` in `.env` file.
Example configuration for local storage:

```bash
LOCAL_STORAGE_PATH=/tmp
STORAGE_TYPE=local
```

## Usage

### Running

To run the microservice execute `make run` command in project _file-transfer_ directory.

### Testing

To run tests execute `make test` command in project _file-transfer_ directory.

### Example usage

To upload file you need to send POST request to `/files` endpoint with file in body.
Example request using `curl`:

```bash
curl -X POST "http://localhost:8080/file" \
    -H "Content-Type: application/json" \
    -d '{
        "file_name": "test.txt",
        "user_id": "123"
    }'
```

In response you will receive file id which you can use to download file.

To download file you need to send GET request to `/files/{file_id}` endpoint.
Example request using `curl`:

```bash
curl -X GET "http://localhost:8080/file/123" \
    -H "Content-Type: application/json"
```

### Swagger

To access Swagger documentation go to `http://localhost:8080/swagger/index.html`.
