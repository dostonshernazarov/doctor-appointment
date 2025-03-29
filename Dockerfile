FROM golang:1.21-alpine3.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o main cmd/app/main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8070

CMD ["/app/main"]