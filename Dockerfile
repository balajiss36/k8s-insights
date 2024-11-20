FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

COPY go.mod go.sum ./
RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o /app/main .

FROM alpine:latest

LABEL author="Balaji Shettigar"

WORKDIR /app

COPY --from=builder /app/main .

COPY config.env config.env

EXPOSE 9050

CMD [ "/app/main" ]