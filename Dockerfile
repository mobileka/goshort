FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o goshort ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/goshort .
COPY --from=builder /app/ui /root/ui
CMD ["./goshort"]