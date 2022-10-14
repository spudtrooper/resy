package cli

import (
	"context"

	"github.com/spudtrooper/minimalcli/handler"
	flag "github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/resy/api"
	"github.com/spudtrooper/resy/handlers"
)

func Main(ctx context.Context) error {
	flag.String("term", "", "search term")
	flag.Bool("debug_body", false, "debug body")
	flag.Bool("debug_payload", false, "debug payload")
	flag.String("url_slug", "", "url slug")
	flag.String("location", "", "location")
	flag.Int("venue_id", 0, "venue ID")
	flag.String("start_date", "", "start date")
	flag.String("end_date", "", "end date")

	client, err := api.NewClientFromFlags()
	if err != nil {
		return err
	}

	return handler.RunCLI(ctx, handlers.CreateHandlers(client)...)
}
