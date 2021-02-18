FROM golang:1.5-alpine

RUN apk --no-cache add ca-certificates

WORKDIR "/go/src/github.com/wcharczuk/echo"

ADD vendor /go/src/github.com/wcharczuk/echo/vendor
ADD main.go /go/src/github.com/wcharczuk/echo/main.go
RUN go install -mod=vendor github.com/wcharczuk/echo

ENTRYPOINT /go/bin/echo
EXPOSE 5000
