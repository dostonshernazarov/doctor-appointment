FROM golang:1.24.1-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o main cmd/app/main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8070

CMD ["/app/main"]