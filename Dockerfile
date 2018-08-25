FROM golang

ADD . /go/src/github.com/josefigueredo/bootcamp-go-aug18

RUN go get github.com/gorilla/mux

RUN go install github.com/josefigueredo/bootcamp-go-aug18

ENTRYPOINT /go/bin/bootcamp-go-aug18

EXPOSE 4567