FROM golang:1.21-alpine

COPY .env /app/.env

WORKDIR /app

RUN go mod tidy

RUN go build -o app .

CMD ["/app/app"]
