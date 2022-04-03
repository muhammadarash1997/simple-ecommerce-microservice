FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o binary-cart-microservice

CMD ./binary-cart-microservice