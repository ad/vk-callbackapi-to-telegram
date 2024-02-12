FROM --platform=${BUILDPLATFORM:-linux/amd64} danielapatin/homeassistant-addon-golang-template as builder

ARG BUILD_VERSION
ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR $GOPATH/src/app
COPY . .
COPY config.json /config.json
RUN echo "Building for ${TARGETOS}/${TARGETARCH} with version ${BUILD_VERSION}"
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s -X main.version=${BUILD_VERSION}" -o /go/bin/app main.go

FROM --platform=${TARGETPLATFORM:-linux/amd64} scratch
WORKDIR /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY config.json /config.json
COPY --from=builder /go/bin/app /go/bin/app
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
