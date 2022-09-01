FROM golang:1.19-alpine3.16

WORKDIR /go/src/api

RUN go install github.com/gofiber/cli/fiber@latest

COPY . .

RUN go mod download

CMD ["fiber", "dev"]