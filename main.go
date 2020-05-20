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
	if len(args) > 2 {
		log.Fatalf("Port appears to have not been specificed. Example: %v", "us.pool.ntp.org:123")
	}
	time, err := ntp.GetTime(ntp.Options{Server: args[0], Port: args[1]})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Retrieved time from %s", host)
	fmt.Printf("%v\n", time)

}
