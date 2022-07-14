package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/deepsourcelabs/demo-go/pkg/cache"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func (c routerCtx) Logger() http.Handler {
	return handlers.CombinedLoggingHandler(os.Stderr, c.Router)
}

type routerCtx struct {
	ctx context.Context
	*mux.Router
	db cache.CacheManager
}

func (c *routerCtx) CacheRegister() {
	c.db = cache.NewDB()
}

func NewServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handler,
	}
}

func signalHandler() (<-chan os.Signal, func()) {
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return sig, func() { close(sig) }
}
