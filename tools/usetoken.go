package main

import (
	"flag"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/resy/api"
)

func useCookie() error {
	token := flag.Arg(0)
	creds, err := api.ReadCredsFromFlags()
	if err != nil {
		return errors.Errorf("api.ReadCredsFromFlags: %v", err)
	}
	creds.Token = token
	if err := api.WriteCredsFromFlags(creds); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	check.Err(useCookie())
}
