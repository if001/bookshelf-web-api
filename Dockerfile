FROM golang:1.12 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/bookshelf-web-api
COPY . .

RUN \
go get github.com/go-sql-driver/mysql &&\
go get github.com/jinzhu/gorm &&\
go get github.com/jinzhu/gorm/dialects/mysql &&\
github.com/julienschmidt/httprouter &&\

RUN go build cmd/api/main.go -o app



FROM alpine:3.9
RUN apk add --no-cache --virtual build-dependencies gcc make autoconf libc-dev libtool &&\
    apk add --no-cache imagemagick
COPY --from=builder /go/src/bookshelf-web-api/app /go_app/app
COPY --from=builder /go/src/ImageConvert/sphericalpano2rect /go_app/sphericalpano2rect
RUN mkdir -p /go_app/img
ENTRYPOINT ["/go_app/app]