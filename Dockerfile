FROM golang:latest
LABEL maintainer="kurizcan"
WORKDIR $GOPATH/src/experiment
ADD . $GOPATH/src/experiment
RUN go build .
EXPOSE 8081
ENTRYPOINT ["./experiment"]