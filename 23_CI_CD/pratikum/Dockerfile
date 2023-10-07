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
EXPOSE 8000

# Wait for MySQL service
CMD ["./server"] 