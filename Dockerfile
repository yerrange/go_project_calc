FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o calc_server ./cmd/server

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/calc_server .

EXPOSE 8080 50051

CMD ["./calc_server"]
