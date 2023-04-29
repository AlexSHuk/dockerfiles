# Сборка исполняемого файла
FROM golang:1.16-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . .
RUN go get -d -v
RUN go build -o app

# Создание образа
FROM alpine:3.13
#LABEL maintainer="Your Name <your.email@example.com>"

RUN apk update && apk add --no-cache ca-certificates

WORKDIR /root/
COPY --from=builder /app/app .

EXPOSE 80

CMD ["./app"]
