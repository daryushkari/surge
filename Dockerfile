# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

EXPOSE 8080

COPY config.json ./
COPY test-config.json ./
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /surge
CMD ["/surge"]