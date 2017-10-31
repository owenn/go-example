FROM golang:1.9

USER nobody

RUN mkdir -p /go/src/github.com/openshift/go-example
WORKDIR /go/src/github.com/openshift/go-example

COPY . /go/src/github.com/openshift/go-example

RUN go-wrapper download && go-wrapper install

CMD ["go-wrapper", "run"
