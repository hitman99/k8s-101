FROM golang:1.12-alpine3.9

COPY . /go/src/github.com/hitman99/k8s-101/demo-app
WORKDIR /go/src/github.com/hitman99/k8s-101/demo-app
RUN go build -o demoapp

FROM alpine:3.9
LABEL maintainer="tomas@adomavicius.com"

RUN apk --no-cache add ca-certificates
WORKDIR /kubernetes-101
COPY --from=0 /go/src/github.com/hitman99/k8s-101/demo-app/demoapp /kubernetes-101/demoapp
COPY static /kubernetes-101/static
ENV PATH="/kubernetes-101/:${PATH}"

CMD ["demoapp"]