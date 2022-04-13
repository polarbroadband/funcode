package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	// api listen TCP port
	PORT = "8089"
)

func init() {
	// config package level default logger
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
}

func main() {
	log.Infof("API server start, listen on port %s", PORT)
	http.HandleFunc("/", greeting)

	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatalf("Server listening error: %+v", err)
	}
}
