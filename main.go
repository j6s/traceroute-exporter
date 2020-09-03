package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type Hop struct {
	Number                 int
	Address                string
	Hostname               string
	AverageRoundTripTimeMs float64
}

type TraceRouteResult struct {
	Destination  string
	Hops         []Hop
	NumberOfHops int
}

var (
	roundTripRegex    = regexp.MustCompile(`(\d+\.\d+)\s+ms`)
	ipExtractionRegex = regexp.MustCompile(`\((.*)\)`)
	listen            = flag.String("listen", ":9094", "Port (and host) to listen on")
)

func traceroute(host string) (TraceRouteResult, error) {

	result := TraceRouteResult{Destination: host}
	bytes, err := exec.Command("traceroute", host).Output()
	if err != nil {
		return result, err
	}

	lines := strings.Split(string(bytes), "\n")
	for i, line := range lines[1 : len(lines)-1] {
		parts := regexp.MustCompile("\\s+").Split(line, -1)
		if parts[2] == "*" {
			continue
		}

		roundTripTimeSum := 0.0
		roundTripTimeCount := 0.0
		for _, timeString := range roundTripRegex.FindAllString(line, -1) {
			time, err := strconv.ParseFloat(roundTripRegex.ReplaceAllString(timeString, "$1"), 64)
			if err != nil {
				return result, err
			}
			roundTripTimeSum += time
			roundTripTimeCount++
		}

		result.Hops = append(result.Hops, Hop{
			Number:                 i,
			Hostname:               parts[2],
			Address:                ipExtractionRegex.ReplaceAllString(parts[3], "$1"),
			AverageRoundTripTimeMs: roundTripTimeSum / roundTripTimeCount,
		})
		result.NumberOfHops = i
	}

	return result, nil
}

func format(result TraceRouteResult) string {
	lines := []string{
		fmt.Sprintf("num_hops{destination=\"%s\"} %d", result.Destination, result.NumberOfHops),
	}

	for _, hop := range result.Hops {
		lines = append(lines, fmt.Sprintf(
			"hop_roundtrip_time_ms{destination=\"%s\",hop=\"%d\",hostname=\"%s\",ip=\"%s\"} %f",
			result.Destination,
			hop.Number,
			hop.Hostname,
			hop.Address,
			hop.AverageRoundTripTimeMs,
		))
	}

	return strings.Join(lines, "\n")
}

func main() {
	flag.Parse()
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		for _, destination := range r.URL.Query()["destination"] {
			result, err := traceroute(destination)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintf(w, "%s\n", format(result))
		}
	})

	log.Printf("Starting to listen on %s", *listen)
	err := http.ListenAndServe(*listen, nil)
	if err != nil {
		log.Fatal(err)
	}
}
