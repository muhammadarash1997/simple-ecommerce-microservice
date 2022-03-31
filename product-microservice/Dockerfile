FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o binary-product-microservice

EXPOSE 8081

CMD ./binary-product-microservice