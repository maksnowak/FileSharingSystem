FROM golang:1.23-alpine
LABEL authors="Paweł (Duszke333) Wysocki"
COPY . .
RUN go build
ENTRYPOINT ["./accounts"]
