package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/kyuff/anchor"
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
	server := &http.Server{}
	server.Addr = ":8080"
	return &HttpServer{
		server: server,
	}, nil
})
