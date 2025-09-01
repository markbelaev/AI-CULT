from golang:1.25 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tg_bot ./cmd/bot
FROM alpine:3.22.1
WORKDIR /app
COPY --from=builder /app/tg_bot .
EXPOSE 8080
CMD ["./tg_bot"]
