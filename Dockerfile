
FROM golang

ADD . github.com/Nepooomuk/golang-knative-helloworld

RUN go install github.com/Nepooomuk/golang-knative-helloworld

ENTRYPOINT /go/bin/golang-knative-helloworld

EXPOSE 8080