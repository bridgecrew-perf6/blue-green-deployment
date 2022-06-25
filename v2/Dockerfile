
FROM golang:1.11

RUN mkdir /go/src/app
ADD . /go/src/app

WORKDIR /go/src/app

RUN go get -u github.com/gorilla/mux
RUN go build -o main .

CMD ["/go/src/app/main"]