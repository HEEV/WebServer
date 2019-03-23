package main

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime"

	"github.com/HEEV/WebServer/chat"

	log "github.com/sirupsen/logrus"
)

// InitializeRoutes initializes all the endpoints for the server
func initializeRoutes(hub *chat.Hub) {
	log.Info("Initializing routes...")

	// Handle all HTTP requests
	http.HandleFunc("/", httpHandler)

	// Handle websocket connections
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})

	// Handle static file requests for javascript
	http.Handle("/static/js/",
		http.StripPrefix("/static/js/",
			http.FileServer(http.Dir("./static/js"))))

	// Handle static file requests for CSS
	http.Handle("/static/css/",
		http.StripPrefix("/static/css/",
			http.FileServer(http.Dir("./static/css"))))

	// Handle static file requests for images
	http.Handle("/static/img/",
		http.StripPrefix("/static/img/",
			http.FileServer(http.Dir("./static/img"))))

	http.HandleFunc("/API", APIHandler)
}

// httpHandler handles all HTTP requests sent to the server
func httpHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		rootHandler(w, r)
		break
	case "/graph":
		graphHandler(w, r)
		break
	default:
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
}

// rootHandler handles requests to the server root
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Values that will be used in the template
	vals := map[string]string{
		"version": runtime.Version(),
	}

	// Parse the root page template
	t, _ := template.ParseFiles("templates/root.html")

	// Respond with the template filled with the values
	t.Execute(w, vals)
}


// graphHandler handles requests to the /graph page
func graphHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Values that will be used in the template
	vals := map[string]string{
		"version": runtime.Version(),
	}

	// Parse the root page template
	t, _ := template.ParseFiles("templates/graph.html")

	// Respond with the template filled with the values
	t.Execute(w, vals)
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	ErrorCheck(w, r, "APIHandler")
}


func GetCarName(w http.ResponseWriter, r *http.Request){
	ErrorCheck(w, r, "APIHandler")
	if r.Method == "POST" && r.URL.Query()["CarId"] != nil {
		var CarId = r.URL.Query()["CarId"]
		//function in model.php
		var carName = GetCarName(CarId)
		fmt.Println(carName);
	}
}

func ErrorCheck(w http.ResponseWriter, r *http.Request, method string){
	if r.URL.Path != "/" {
		http.Error(w, "Not found for " + method,   http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed " + method , http.StatusMethodNotAllowed)
	}
}