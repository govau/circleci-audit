FROM golang:1.11 AS builder

# Install dep
RUN curl -L https://github.com/golang/dep/releases/download/v0.5.1/dep-linux-amd64 > /usr/local/bin/dep && \
  chmod a+x /usr/local/bin/dep

COPY . /go/src/github.com/govau/circleci-audit

WORKDIR /go/src/github.com/govau/circleci-audit

# If we don't disable CGO, the binary won't work. Unsure why?
RUN dep ensure && \
  go test ./... && \
  CGO_ENABLED=0 go install

FROM alpine:3.9

RUN apk add --update \
  bash \
  curl \
  git \
  jq \
  && \
  rm -rf /var/cache/apk/*

COPY --from=builder /go/bin/circleci-audit /usr/bin/circleci-audit

ENTRYPOINT [ "/usr/bin/circleci-audit"]
