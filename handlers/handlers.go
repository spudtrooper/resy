// Package handlers is a bridge between the API and handlers to create a CLI/API server.
package handlers

import (
	"context"
	_ "embed"

	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/resy/api"
)

//go:generate minimalcli gsl --input handlers.go --uri_root "github.com/spudtrooper/resy/blob/main/handlers" --output handlers.go.json
//go:embed handlers.go.json
var SourceLocations []byte

func CreateHandlers(client *api.Client) []handler.Handler {
	b := handler.NewHandlerBuilder()

	b.NewHandler("Search",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.SearchParams)
			return client.Search(p.Term, p.Options()...)
		},
		api.SearchParams{},
	)

	b.NewHandler("Venue",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.VenueParams)
			return client.Venue(p.UrlSlug, p.Location, p.Options()...)
		},
		api.VenueParams{},
	)

	b.NewHandler("Config",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.ConfigParams)
			return client.Config(p.VenueID, p.Options()...)
		},
		api.ConfigParams{},
	)

	b.NewHandler("Calendar",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.CalendarParams)
			return client.Calendar(p.VenueID, p.Options()...)
		},
		api.CalendarParams{},
	)

	return b.Build()
}
