FROM golang:1.10.0-alpine3.7 as builder
ENV buildpath=/go/src/github.com/{{GITHUB_USERNAME}}/{{GITHUB_REPO}}
RUN mkdir -p $buildpath
RUN apk add --update make bash git
WORKDIR $buildpath

COPY . .

RUN make build
RUN go get -u golang.org/x/tools/cmd/goimports
RUN make test


FROM alpine:3.7
RUN apk --no-cache add ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/{{GITHUB_USERNAME}}/{{GITHUB_REPO}}/_release/{{GITHUB_REPO}} /{{GITHUB_REPO}}


ENTRYPOINT ["/{{GITHUB_REPO}}"]
