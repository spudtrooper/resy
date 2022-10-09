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
	mux, err := handler.CreateHandler(ctx, handlers,
		handler.CreateHandlerPrefix("api"),
		handler.CreateHandlerIndexTitle("unofficial resy API"),
		handler.CreateHandlerFooterHTML(`Details: <a target="_" href="https://github.com/spudtrooper/resy">github.com/spudtrooper/resy</a>`),
		handler.CreateHandlerSourceLinks(true),
		handler.CreateHandlerHandlersFiles([]string{"handlers/handlers.go"}),
		handler.CreateHandlerSourceLinkURIRoot("https://github.com/spudtrooper/resy/blob/main"),
	)
	if err != nil {
		return err
	}
	mux.Handle("/", http.RedirectHandler("/api", http.StatusSeeOther))

	log.Printf("listening on %s", hostPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		return err
	}

	return nil
}
