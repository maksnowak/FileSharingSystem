FROM golang:1.23-alpine
LABEL authors="Piotr (Depermitto) Jabłoński"
COPY . .
RUN go build
ENTRYPOINT ["./file-transfer"]