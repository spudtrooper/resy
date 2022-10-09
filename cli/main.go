package cli

import (
	"context"

	"github.com/spudtrooper/goutil/flags"
	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/resy/api"
	"github.com/spudtrooper/resy/handlers"
)

var (
	term         = flags.String("term", "search term")
	debugBody    = flags.Bool("debug_body", "debug body")
	debugPayload = flags.Bool("debug_payload", "debug payload")
	urlSlug      = flags.String("url_slug", "url slug")
	location     = flags.String("location", "location")
)

func Main(ctx context.Context) error {
	adp := handler.NewCLIAdapter()
	adp.BindStringFlag("term", term)
	adp.BindBoolFlag("debug_body", debugBody)
	adp.BindBoolFlag("debug_payload", debugPayload)
	adp.BindStringFlag("url_slug", urlSlug)
	adp.BindStringFlag("location", location)

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
