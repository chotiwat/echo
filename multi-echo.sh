#!/bin/sh

env PORT=6000 /go/bin/echo &
env PORT=7000 /go/bin/echo &

/go/bin/echo