FROM golang:1.8.4-jessie as builder
ENV buildpath=/usr/local/go/src/build/{{GITHUB_REPO}}
ARG build=notSet
RUN mkdir -p $buildpath
ADD . $buildpath
WORKDIR $buildpath

RUN make build/release

FROM alpine:3.6
COPY --from=builder /usr/local/go/src/build/{{GITHUB_REPO}}/_release/{{GITHUB_REPO}} /{{GITHUB_REPO}}

ENTRYPOINT ["/{{GITHUB_REPO}}"]
