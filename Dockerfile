FROM golang
MAINTAINER Roy van de Water <roy.v.water@gmail.com>
EXPOSE 80

ADD https://raw.githubusercontent.com/pote/gpm/v1.3.2/bin/gpm /go/bin/
RUN chmod +x /go/bin/gpm

COPY Godeps /go/src/github.com/royvandewater/gogoslow4me/
WORKDIR /go/src/github.com/royvandewater/gogoslow4me
RUN gpm install

COPY . /go/src/github.com/royvandewater/gogoslow4me

RUN env CGO_ENABLED=0 go build -a -ldflags '-s' .

CMD ["./gogoslow4me"]
