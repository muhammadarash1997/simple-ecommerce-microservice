FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o binary-payment-microservice

CMD ./binary-payment-microservice