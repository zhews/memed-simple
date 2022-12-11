FROM golang:1.19-alpine as build
WORKDIR /build
COPY . .
RUN go build ./cmd/meme
FROM alpine:3.17 as run
WORKDIR /run
COPY --from=build /build/meme meme
RUN addgroup -S memed && adduser -S memed -G memed
USER memed
HEALTHCHECK CMD wget --no-verbose --tries=1 --spider "http://localhost:$MEMED_PORT/health" || exit 1
CMD ["/run/meme"]