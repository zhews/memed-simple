FROM golang:1.19-alpine as build
WORKDIR /build
COPY . .
RUN go build ./cmd/meme
FROM alpine:3.17 as run
WORKDIR /run
RUN addgroup -S memed && adduser -S memed -G memed
COPY --from=build --chown=memed:memed /build/meme meme
RUN mkdir static && chown memed:memed static
VOLUME ["/run/static"]
USER memed
HEALTHCHECK --interval=5s CMD wget --no-verbose --tries=1 --spider "http://localhost:$MEMED_PORT/health" || exit 1
CMD ["/run/meme"]