package frontend

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/resy/api"
	"github.com/spudtrooper/resy/handlers"
)

func ListenAndServe(ctx context.Context, client *api.Client, port int, host string) error {
	var hostPort string
	if host == "localhost" {
		hostPort = fmt.Sprintf("http://localhost:%d", port)
	} else {
		hostPort = fmt.Sprintf("https://%s", host)
	}

	handlers := handlers.CreateHandlers(client)
	mux := http.NewServeMux()
	if err := handler.AddHandlers(ctx, mux, handlers,
		handler.AddHandlersPrefix("api"),
		handler.AddHandlersIndexTitle("unofficial resy API"),
		handler.AddHandlersFooterHTML(`Details: <a target="_" href="//github.com/spudtrooper/resy">github.com/spudtrooper/resy</a>`),
		handler.AddHandlersSourceLinks(true),
		handler.AddHandlersHandlersFiles([]string{"handlers/handlers.go"}),
		handler.AddHandlersSourceLinkURIRoot("//github.com/spudtrooper/resy/blob/main"),
	); err != nil {
		return err
	}
	mux.Handle("/", http.RedirectHandler("/api", http.StatusSeeOther))

	log.Printf("listening on %s", hostPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		return err
	}

	return nil
}
