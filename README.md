# Prometheus Traceroute exporter

This exporter for [prometheus](prometheus.io/) exposes traceroute information in
order to monitor the number of hops are needed to a given destination host.

## Usage

```
# Directly
$ go run main.go --listen ':9094'

# Docker
$ docker run --rm thej6s/traceroute-exporter --listen ':9094'
```

Now call `localhost:9094/metrics?destination=1.1.1.1`.
For multiple destinations, add the `destination` query parameter multiple times.

**Note:** Using the docker contains will always add an additional hop for the machine-internal
routing between the docker client and host.

## Example metrics

Called using `http://localhost:9094/metrics?destination=1.1.1.1&destination=8.8.8.8`

```
num_hops{destination="1.1.1.1"} 6
hop_roundtrip_time_ms{destination="1.1.1.1",hop="0",hostname="172.17.0.1",ip="172.17.0.1"} 0.025667
hop_roundtrip_time_ms{destination="1.1.1.1",hop="1",hostname="10.0.0.1",ip="10.0.0.1"} 0.587333
hop_roundtrip_time_ms{destination="1.1.1.1",hop="2",hostname="217.147.60.88",ip="217.147.60.88"} 6.153000
hop_roundtrip_time_ms{destination="1.1.1.1",hop="3",hostname="ae0-405.bas10.core-backbone.com",ip="5.56.19.1"} 7.243000
hop_roundtrip_time_ms{destination="1.1.1.1",hop="4",hostname="ae1-2053.prg10.core-backbone.com",ip="81.95.15.118"} 18.256000
hop_roundtrip_time_ms{destination="1.1.1.1",hop="5",hostname="nix4.cloudflare.com",ip="91.210.16.171"} 23.014333
hop_roundtrip_time_ms{destination="1.1.1.1",hop="6",hostname="one.one.one.one",ip="1.1.1.1"} 19.055333
num_hops{destination="8.8.8.8"} 7
hop_roundtrip_time_ms{destination="8.8.8.8",hop="0",hostname="172.17.0.1",ip="172.17.0.1"} 0.444000
hop_roundtrip_time_ms{destination="8.8.8.8",hop="1",hostname="10.0.0.1",ip="10.0.0.1"} 1.152000
hop_roundtrip_time_ms{destination="8.8.8.8",hop="2",hostname="217.147.60.88",ip="217.147.60.88"} 6.687000
hop_roundtrip_time_ms{destination="8.8.8.8",hop="3",hostname="217.147.60.125",ip="217.147.60.125"} 11.651333
hop_roundtrip_time_ms{destination="8.8.8.8",hop="4",hostname="sis-googl.stiegeler.com",ip="217.147.60.105"} 12.905333
hop_roundtrip_time_ms{destination="8.8.8.8",hop="5",hostname="108.170.251.193",ip="108.170.251.193"} 10.637000
hop_roundtrip_time_ms{destination="8.8.8.8",hop="6",hostname="108.170.235.245",ip="108.170.235.245"} 10.165333
hop_roundtrip_time_ms{destination="8.8.8.8",hop="7",hostname="dns.google",ip="8.8.8.8"} 10.542667
```