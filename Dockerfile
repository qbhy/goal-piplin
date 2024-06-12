FROM golang:1.20 as builder
LABEL maintainer="qbhy <qbhy0715@qq.com>"

WORKDIR /app

COPY . /app
ENV CGO_ENABLED=0
ENV GOOS=linux
#ENV GOARCH=amd64
#ENV GOPROXY=https://proxy.golang.com.cn,direct
RUN go build -ldflags="-s -w" -o piplin main.go

FROM alpine

RUN apk add git openssh-client
RUN mkdir -p ~/.ssh
RUN ssh-keyscan github.com > ~/.ssh/known_hosts

WORKDIR /var/www
COPY --from=builder /app/piplin /var/www/piplin
COPY env.toml .
COPY entrypoint.sh .
COPY database database

# run
ENTRYPOINT ["./entrypoint.sh"]