FROM golang:1.14-alpine

RUN apk update
RUN apk add git vim redis make

WORKDIR /go/src/github.com/aoffy-kku/minemind-backend

ADD . .

RUN go build main.go
EXPOSE 1321
CMD ["./main"]

