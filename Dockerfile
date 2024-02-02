# STAGE 1: Build the binary
FROM golang:alpine as build

# Build arguments
ARG BUILD_ARCH
ARG BUILD_VERSION

# Install git for go get command and get the repo
RUN \
    apk update && apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR $GOPATH/src/app

COPY . .
COPY config.json /config.json

# Install dependencies and build the binary to target platform
RUN \
    if [ "${BUILD_ARCH}" = "armhf" ]; then \
    GOOS=linux GOARCH=arm go build -o /go/bin/app; \
    elif [ "${BUILD_ARCH}" = "armv7" ]; then \
    GOOS=linux GOARM=7 GOARCH=arm go build -o /go/bin/app; \
    elif [ "${BUILD_ARCH}" = "aarch64" ]; then \
    GOOS=linux GOARCH=arm64 go build -o /go/bin/app; \
    elif [ "${BUILD_ARCH}" = "i386" ]; then \
    GOOS=linux GOARCH=386 go build -o /go/bin/app; \
    elif [ "${BUILD_ARCH}" = "amd64" ]; then \
    GOOS=linux GOARCH=amd64 go build -o /go/bin/app; \
    else \
    echo 'NOT VALID BUILD'; exit 1; \
    fi

# STAGE 2: Include binary in target add-on container
FROM scratch AS runtime

# Copy binary and the config from build container
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group
COPY --from=build /go/bin/app /go/bin/app
COPY --from=build /config.json /config.json

ENTRYPOINT ["/go/bin/app"]

#
# LABEL target docker image
#

# Build arguments
ARG BUILD_ARCH
ARG BUILD_DATE
ARG BUILD_REF
ARG BUILD_VERSION

# Labels
LABEL \
    io.hass.name="vk-callbackapi-to-telegram" \
    io.hass.description="vk-callbackapi-to-telegram" \
    io.hass.arch="${BUILD_ARCH}" \
    io.hass.version=${BUILD_VERSION} \
    io.hass.type="addon" \
    maintainer="ad <github@apatin.ru>" \
    org.label-schema.description="vk-callbackapi-to-telegram" \
    org.label-schema.build-date=${BUILD_DATE} \
    org.label-schema.name="vk-callbackapi-to-telegram" \
    org.label-schema.schema-version="1.0" \
    org.label-schema.usage="https://gitlab.com/ad/vk-callbackapi-to-telegram/-/blob/master/README.md" \
    org.label-schema.vcs-ref=${BUILD_REF} \
    org.label-schema.vcs-url="https://github.com/ad/vk-callbackapi-to-telegram/" \
    org.label-schema.vendor="HomeAssistant add-ons by ad"
