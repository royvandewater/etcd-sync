FROM golang
MAINTAINER Roy van de Water <roy.v.water@gmail.com>
EXPOSE 80

WORKDIR /go/src/github.com/royvandewater/etcdsync
COPY . /go/src/github.com/royvandewater/etcdsync

RUN env CGO_ENABLED=0 go build -a -ldflags '-s' .

CMD ["./etcdsync"]
