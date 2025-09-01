from golang:1.25 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .
FROM alpine:3.22.1
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
