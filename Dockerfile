FROM golang:alpine
ADD . /go/src/github.com/ryane/k8s-deployment-demo
RUN go install github.com/ryane/k8s-deployment-demo
ENTRYPOINT /go/bin/k8s-deployment-demo
