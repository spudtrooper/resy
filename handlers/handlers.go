// Package handlers is a bridge between the API and handlers to create a CLI/API server.
package handlers

import (
	"context"

	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/resy/api"
)

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

	return b.Build()
}
