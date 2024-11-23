FROM golang:1.23-alpine AS build-stage
LABEL authors="Piotr (Depermitto) Jabłoński"
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o main

FROM scratch
WORKDIR /app
COPY --from=build-stage /build/main .
ENTRYPOINT ["./main"]