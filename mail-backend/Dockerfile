# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD . /app
COPY *.go ./

RUN go build -o /mail-backend

EXPOSE 8080

CMD [ "/mail-backend" ]