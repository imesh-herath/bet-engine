package server

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var PprofPort = 6060

func Profiling() {
	go func() {

		err := http.ListenAndServe(fmt.Sprintf(":%d", PprofPort), nil)
		if err != nil {
			log.Panicf("PProf server on %d cannot start!", PprofPort)
		}
	}()
}
