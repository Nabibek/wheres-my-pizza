FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go build -o /app/cmd/main ./cmd/main.go

CMD ["./myapp --mode=order"]
