FROM registry.sandbox.k8s.centrio.com/proxy:2d24d26

RUN apk --no-cache add ca-certificates

WORKDIR "/go/src/github.com/wcharczuk/echo"

ADD main.go /go/src/github.com/wcharczuk/echo/main.go
ADD vendor /go/src/github.com/wcharczuk/echo/vendor
ADD start.sh /go/src/github.com/wcharczuk/echo/start.sh

RUN go install github.com/wcharczuk/echo

ENTRYPOINT ["./start.sh"]
