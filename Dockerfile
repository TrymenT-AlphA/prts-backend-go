FROM golang:1.18-alpine

ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./dist/prts-backend

WORKDIR /build/dist

RUN apt-get -q update && apt-get -qy install netcat

CMD [ "sh", "-c", "", "./prts-backend" ]
