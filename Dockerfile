# dynamic config
ARG             BUILD_ARG_VERSION=latest

# build
FROM            golang:1.16.6-alpine as builder
ARG             BUILD_ARG_VERSION
RUN             apk add --no-cache git gcc musl-dev make
ENV             GO111MODULE=on
ENV             VERSION=${BUILD_ARG_VERSION}

WORKDIR         /src
COPY            go.* ./
RUN             go mod download
COPY            . ./
RUN             make vendor
RUN             make build

# minimalist runtime
FROM            alpine:3.13
ARG             BUILD_ARG_VERSION
LABEL   org.label-schema.name="line-coverage" \
    org.label-schema.description="Analysed coverage profile file from Golang code to calculate and report line coverage result" \
    org.label-schema.version=$BUILD_ARG_VERSION \
    org.label-schema.vcs-url="https://github.com/Liu-Chunhui/line-coverage" \
    org.label-schema.help="docker run -it --rm yesino/line-coverage"

COPY            --from=builder /src/bin/line-coverage /bin/

ENTRYPOINT ["/bin/line-coverage"]

CMD [ "--help" ]
