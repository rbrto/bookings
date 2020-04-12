FROM golang:1.14-alpine as builder

LABEL maintainer="robertoesparza3@gmail.com"

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -a -installsuffix cgo -o bookings .


FROM alpine:3.9.2

WORKDIR /root/

COPY --from=builder /app .

EXPOSE 8080

ENTRYPOINT ["./bookings"]
