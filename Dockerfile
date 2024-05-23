FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o ./cmd/main ./cmd/cmd.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/cmd /app/cmd
CMD ["./cmd/main"]