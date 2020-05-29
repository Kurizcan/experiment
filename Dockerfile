FROM golang:latest
LABEL maintainer="kurizcan"
ENV GO111MODULE=on GOPROXY=https://goproxy.io,direct
WORKDIR $GOPATH/src/experiment
ADD . $GOPATH/src/experiment
RUN go build .
EXPOSE 8081
ENTRYPOINT ["./experiment"]