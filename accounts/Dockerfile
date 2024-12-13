FROM golang:1.23-alpine AS build-stage
LABEL authors="Paweł (Duszke333) Wysocki"
WORKDIR /build
COPY . .
COPY ../.env .
RUN go mod download
RUN go build -o main

FROM scratch AS production-stage
WORKDIR /app
COPY --from=build-stage /build/main .
COPY --from=build-stage /build/.env .
ENTRYPOINT ["./main"]