package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/kyuff/anchor"
	"github.com/kyuff/example-petstore/generated/api"
	"github.com/kyuff/example-petstore/internal/petstore"
)

type HttpServer struct {
	server *http.Server
}

func (h *HttpServer) Start(ctx context.Context) error {
	h.server.BaseContext = func(_ net.Listener) context.Context {
		return ctx
	}
	err := h.server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	if err != nil {
		slog.Error(fmt.Sprintf("http server start error %T: %v", err, err))
	}

	return err
}

func (h *HttpServer) Close() error {
	return h.server.Shutdown(context.Background())
}

var httpServer = anchor.Singleton(func() (*HttpServer, error) {
	return &HttpServer{
		server: &http.Server{
			Addr: ":8080",
			Handler: api.HandlerFromMux(
				api.NewStrictHandler(httpApi(), nil),
				http.NewServeMux(),
			),
		},
	}, nil
})

var httpApi = anchor.Singleton(func() (*petstore.API, error) {
	return &petstore.API{}, nil
})
