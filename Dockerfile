
FROM golang:1.22-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o my-app ./cmd

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/my-app .
EXPOSE 8080
CMD ["./my-app"]