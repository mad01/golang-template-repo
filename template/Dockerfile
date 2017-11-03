FROM golang:1.8.4-jessie as builder
ENV buildpath=/go/src/github.com/{{GITHUB_USERNAME}}/{{GITHUB_REPO}}
RUN mkdir -p $buildpath
WORKDIR $buildpath

COPY . .

RUN make build/release

FROM debian:8
COPY --from=builder /go/src/github.com/{{GITHUB_USERNAME}}/{{GITHUB_REPO}}/_release/{{GITHUB_REPO}} /{{GITHUB_REPO}}

ENTRYPOINT ["/{{GITHUB_REPO}}"]
