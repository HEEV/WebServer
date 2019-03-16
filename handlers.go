package main

import (
	"html/template"
	"net/http"

	"runtime"

	"github.com/HEEV/WebServer/chat"

	log "github.com/sirupsen/logrus"
)

// InitializeRoutes initializes all the endpoints for the server
func initializeRoutes(hub *chat.Hub) {
	log.Info("Initializing routes...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})
	http.HandleFunc("/API", APIHandler)
}

// rootHandler handles requests to the server root
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vals := map[string]string{
		"version": runtime.Version(),
	}

	t, _ := template.ParseFiles("templates/root.html")
	t.Execute(w, vals)
}
/*
//TODO:make function that grabs all the data from db and then puts it into API
func APIHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	vals:
}
*/