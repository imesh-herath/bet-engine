package bootstrap

import (
	"bet-engine/http/server"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var httpServer *server.Server

func Init() {
	server.Profiling()

	httpServer = server.NewServer()

	// Start http server
	err := httpServer.Start()
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C), SIGKILL, SIGQUIT or SIGTERM (Ctrl+/)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	// Block until we receive our signal
	signal := <-c
	log.Println("bootstrap.init.Start", fmt.Sprintf("Received Signal: %s", signal))

	// Start destructing the process
	destruct()
}

func destruct() {
	httpServer.Stop()
}
