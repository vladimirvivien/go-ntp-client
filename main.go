package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/vladimirvivien/go-ntp-client/ntp"
)

// This program implements a trivial NTP client over UDP.
//
// Usage:
// time -e <host endpoint as addr:port>
//
func main() {
	var host string
	flag.StringVar(&host, "e", "us.pool.ntp.org:123", "NTP host")
	flag.Parse()

	args := strings.SplitN(host, ":", 2)
	c := ntp.Client{Server: args[0], Port: args[1]}

	time, err := c.Time()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Retrieved time from %s", host)
	fmt.Printf("%v\n", time)

}
