FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o binary-apigateway

EXPOSE 8080

CMD ./binary-apigateway