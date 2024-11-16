FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY .env .env

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

COPY .env .env

EXPOSE 9050

CMD [ "/app/main" ]