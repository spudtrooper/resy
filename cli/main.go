package cli

import (
	"context"
	"flag"

	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/resy/api"
	"github.com/spudtrooper/resy/handlers"
)

var (
	term      = flag.String("term", "", "search term")
	debugBody = flag.Bool("debug_body", false, "debug body")
)

func Main(ctx context.Context) error {
	adp := handler.NewCLIAdapter()
	adp.BindStringFlag("term", term)
	adp.BindBoolFlag("debug_body", debugBody)

	client, err := api.NewClientFromFlags()
	if err != nil {
		return err
	}

	handlers := handlers.CreateHandlers(client)
	app := handler.CreateCLI(adp, handlers...)

	if err := app.Run(ctx); err != nil {
		return err
	}

	return nil
}
