FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o binary-payment-microservice

EXPOSE 8082

CMD ./binary-payment-microservice