FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o user-service .

EXPOSE 8080

CMD ["./user-service"]