FROM golang:1.10.1
WORKDIR /go/src/github.com/{{GITHUB_USERNAME}}/{{GITHUB_REPO}}

RUN go get github.com/golang/dep/cmd/dep
RUN go get golang.org/x/tools/cmd/goimports
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -v -vendor-only

COPY . .
RUN CGO_ENABLED=0 GOOS=linux make install
RUN CGO_ENABLED=0 GOOS=linux make test

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/bin/{{GITHUB_REPO}} /bin/{{GITHUB_REPO}}
ENTRYPOINT ["/bin/{{GITHUB_REPO}}"]
