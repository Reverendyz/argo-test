FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

RUN adduser -D -g '' appuser

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

USER appuser

ENV HOST=0.0.0.0

ENV PORT=8080

ENTRYPOINT ["./main"]
