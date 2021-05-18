FROM golang:1.15 as builder

COPY ./ /build
RUN go build -o /traceroute-exporter /build/main.go


FROM alpine:3.13

RUN apk add --no-cache tcptraceroute libc6-compat
COPY --from=builder /traceroute-exporter /usr/local/bin/traceroute-exporter

ENTRYPOINT traceroute-exporter
