package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/resy/api"
	"github.com/spudtrooper/resy/frontend"
)

var (
	portForTesting = flag.Int("port_for_testing", 0, "port to listen on")
	host           = flag.String("host", "unofficial-resy-api.herokuapp.com", "host name")
)

func main() {
	flag.Parse()
	var port int
	if *portForTesting != 0 {
		port = *portForTesting
	} else {
		p, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatalf("invalid port: %v", err)
		}
		port = p
	}
	if port == 0 {
		log.Fatalf("port is required")
	}
	if *host == "" {
		log.Fatalf("host is required")
	}

	ctx := context.Background()
	client := api.NewClient("")
	check.Err(frontend.ListenAndServe(ctx, client, port, *host))
}
