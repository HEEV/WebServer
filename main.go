package main

import (
	"fmt"

	// Standard lib imports
	"flag" // Used for command line flag parsing

	// Handle network
	"net/http"

	// Server imports
	"./chat"

	// Logging imports
	debugFname "github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
)

var port int

func init() {
	var debug, verbose bool
	flag.BoolVar(&debug, "debug", false, "Run the API in debug mode")
	flag.BoolVar(&debug, "d", false, "Run the API in debug mode")
	flag.BoolVar(&verbose, "verbose", false, "Run the API in verbose mode")
	flag.BoolVar(&verbose, "v", false, "Run the API in verbose mode")
	flag.IntVar(&port, "port", 8000, "Specify which port the API will run on")
	flag.Parse()

	log.Info("Initializing server...")

	if debug {
		log.Warn("Starting server in debug mode!")
	} else if verbose {
		log.Warn("Starting server in verbose mode!")
	}

	if debug || verbose {
		log.SetLevel(log.DebugLevel)
		log.AddHook(debugFname.NewHook())
	}
}

func main() {
	// Initialize websocket hub
	hub := chat.NewHub()
	go hub.Run()

	// Initialize HTTP server instance
	initializeRoutes(hub)

	log.Warnf("Server starting on :%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
