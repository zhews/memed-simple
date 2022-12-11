FROM golang:1.19-alpine as build
WORKDIR /build
COPY . .
RUN go build ./cmd/user
FROM alpine:3.17 as run
WORKDIR /run
COPY --from=build /build/user user
RUN addgroup -S memed && adduser -S memed -G memed
USER memed
HEALTHCHECK --interval=5s CMD wget --no-verbose --tries=1 --spider "http://localhost:$MEMED_PORT/health" || exit 1
CMD ["/run/user"]