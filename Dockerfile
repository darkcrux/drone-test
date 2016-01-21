FROM golang:latest

ADD . /echo
WORKDIR /echo
RUN go build .

ENTRYPOINT ["/echo/echo"]
CMD []
