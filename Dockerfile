FROM alpine:3.16.2
RUN apk add --no-cache \
  curl \
  git \
  musl-dev

RUN mkdir /repo && \
  git config --global --add safe.directory /repo

COPY vrsn /
ENTRYPOINT ["/vrsn"]
