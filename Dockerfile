# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

EXPOSE 9000
EXPOSE 8080

COPY config.json ./
COPY config.test-json ./
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /surge
CMD ["/surge"]