FROM golang:1.20.1-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o server

FROM alpine
WORKDIR /root/app
COPY --from=builder /app/. .

ENV NETWORK_NAME=go-database
EXPOSE 8000
CMD ["./server"]