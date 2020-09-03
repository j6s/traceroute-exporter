FROM golang:1.15.1

RUN apt-get update && apt-get install -y traceroute
COPY main.go /traceroute-exporter.go

ENTRYPOINT go run /traceroute-exporter.go