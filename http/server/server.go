package server

import (
	"bet-engine/http/router"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpSrv *http.Server
	wait    time.Duration
}

func NewServer() *Server {
	// initialize the router
	r := router.Init()

	address := "0.0.0.0:8080"
	srv := new(Server)
	// Initialize HTTP server
	srv.httpSrv = &http.Server{
		Addr: address,

		Handler: r,
		// good practice to set timeouts to avoid Slowloris attacks
		WriteTimeout: time.Minute * 2,
		ReadTimeout:  time.Minute * 2,
		IdleTimeout:  time.Minute * 2,
	}

	return srv
}

func (server *Server) Start() error {
	// run HTTP server in a goroutine so that it doesn't block
	go func() {
		err := server.httpSrv.ListenAndServe()
		if err != nil {
			log.Fatal("http.server.Init", err)
			panic("HTTP server shutting down unexpectedly...")
		}
	}()

	log.Println("http.server.Init", fmt.Sprintf("HTTP server listening on %s", server.httpSrv.Addr))

	return nil
}

func (server *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), server.wait)
	defer cancel()

	err := server.httpSrv.Shutdown(ctx)
	if err != nil {
		log.Println("http.server.ShutDown", "Unable to stop HTTP server")
	}
}
