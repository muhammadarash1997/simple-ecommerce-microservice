FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o binary-cart-microservice

EXPOSE 8083

CMD ./binary-cart-microservice