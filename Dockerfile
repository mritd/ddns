FROM golang:1.17-alpine AS builder

ENV SRC_PATH ${GOPATH}/src/github.com/mritd/ddns

WORKDIR ${SRC_PATH}

COPY . .

RUN set -ex \
    && apk add git \
    && export COMMIT_ID=$(git rev-parse HEAD) \
    && go install -ldflags "-w -s -X 'main.commitID=${COMMIT_ID}'"

FROM alpine

ARG TZ="Asia/Shanghai"

ENV TZ ${TZ}
ENV LANG en_US.UTF-8
ENV LC_ALL en_US.UTF-8
ENV LANGUAGE en_US:en

RUN set -ex \
    && apk add bash tzdata ca-certificates \
    && ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && rm -rf /var/cache/apk/*

COPY --from=builder /go/bin/ddns /usr/local/bin/ddns

ENTRYPOINT ["ddns"]

CMD ["--help"]
