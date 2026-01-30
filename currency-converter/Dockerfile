# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod main.go ./
RUN go build -o server .

# Run stage
FROM alpine:3.19

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server .
COPY index.html .

EXPOSE 8080

ENV PORT=8080

CMD ["./server"]
