FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go build -o myapp

FROM debian:bullseye
WORKDIR /app
COPY --from=builder /app/myapp .
CMD ["./myapp"]
