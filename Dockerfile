FROM registry.sandbox.k8s.centrio.com/proxy:latest

RUN apk --no-cache add ca-certificates

WORKDIR "/go/src/github.com/wcharczuk/echo"

ADD main.go /go/src/github.com/wcharczuk/echo/main.go
ADD vendor /go/src/github.com/wcharczuk/echo/vendor
RUN go install github.com/wcharczuk/echo

RUN env PORT=5000 /go/bin/echo &

ENTRYPOINT /go/bin/proxy
