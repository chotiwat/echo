FROM golang:1.15.5-alpine3.12

RUN apk --no-cache add ca-certificates

WORKDIR "/go/src/github.com/wcharczuk/echo"

ADD vendor /go/src/github.com/wcharczuk/echo/vendor
ADD main.go /go/src/github.com/wcharczuk/echo/main.go
RUN go install github.com/wcharczuk/echo

ENTRYPOINT /go/bin/echo
EXPOSE 5000
