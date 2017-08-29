FROM golang:1.8-alpine

RUN apk --no-cache add ca-certificates

WORKDIR "/go/src/github.com/wcharczuk/echo"

ADD main.go /go/src/github.com/wcharczuk/echo/main.go
ADD vendor /go/src/github.com/wcharczuk/echo/vendor
ADD multi-echo.sh /go/src/github.com/wcharczuk/echo/multi-echo.sh
RUN go install github.com/wcharczuk/echo

ENTRYPOINT ["./multi-echo.sh"]
EXPOSE 5000
