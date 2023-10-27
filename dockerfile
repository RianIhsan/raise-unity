FROM golang:1.21-alpine

# Mengatur variabel lingkungan
ENV DB_HOST=${DB_HOST} \
    DB_PORT=${DB_PORT} \
    DB_USER=${DB_USER} \
    DB_PASSWORD=${DB_PASSWORD} \
    DB_NAME=${DB_NAME} \
    SECRET_KEY=${SECRET_KEY} \
    CLOUDINARY_URL=${CLOUDINARY_URL} \
    SMTP_PASS=${SMTP_PASS} \
    SMTP_USER=${SMTP_USER} \
    SMTP_PORT=${SMTP_PORT}

COPY . /app

WORKDIR /app

RUN go mod tidy

RUN go build -o app .

CMD ["/app/app"]
