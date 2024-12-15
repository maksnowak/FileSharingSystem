# File Transfer

## Description

File transfer microservice is responsible for handling CRUD operations of user files.
It communicates with other services to store retrieve specific files.

## Prerequisites

- [Golang version >=1.23](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Usage

This microservice is dependent on MongoDB database so before starting the web server you need to provide connection configuration to MongoDB database.

### Running MongoDB locally

You can run MongoDB locally with `docker compose` by using `docker compose run mongodb` command in project _file-transfer_ directory.
Then you need to create `.env` file with MongoDB configuration.
Example configuration:

```bash
MONGODB_URI=mongodb://localhost:27017/
MONGODB_DB_USER=root
MONGODB_DB_PASSWORD=example
```

### Connecting to existing instance

Provide connection details in `.env` file.

### Building and downloading packages

```bash
go mod download
go mod build -o main
```

### Running the project

```bash
./main
```

It should start web server on specified port and connect to MongoDB.
