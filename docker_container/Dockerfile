FROM golang

RUN mkdir -p /opt/src
ADD . /opt/src/compress/

ENV GOPATH "$GOPATH:/opt"

WORKDIR /opt/src/compress

# Install MongoDB driver and extra networking package
RUN go get gopkg.in/mgo.v2 net

# Install testify
RUN go get github.com/stretchr/testify

# Install go2xunit
RUN go get github.com/tebeka/go2xunit

# Install logrus and hooks
RUN go get github.com/Sirupsen/logrus
RUN go get github.com/rifflock/lfshook
RUN go get github.com/evalphobia/logrus_sentry

# Install go-yaml for configurations
RUN go get gopkg.in/yaml.v2

# Install cast
RUN go get github.com/spf13/cast
