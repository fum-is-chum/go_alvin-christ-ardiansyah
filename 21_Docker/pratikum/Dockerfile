# Build
FROM golang:1.20.1-alpine as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o server

# Final
FROM alpine

WORKDIR /root/app
COPY --from=builder /app/. .

RUN apk add bash curl
RUN curl -o wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh
RUN chmod +x wait-for-it.sh

ENV NETWORK_NAME=go-database
EXPOSE 8000

# Wait for MySQL service
CMD ["./wait-for-it.sh", "-t 0", "db:3306", "--", "./server"]