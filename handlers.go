package main

import (
	"html/template"
	"net/http"
	"runtime"

	"github.com/HEEV/WebServer/api"
	"github.com/HEEV/WebServer/chat"
	"github.com/gorilla/sessions"

	log "github.com/sirupsen/logrus"
)

var (
	key   = []byte("4KI7AXTDH4VACRRK")
	store = sessions.NewCookieStore(key)
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

	// Handle API endpoint requests
	http.HandleFunc("/API", apiHandler)

	// Handle basic authentication
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/auth", authHandler)
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

	// Validate authentication
	// Get storage cookie store
	session, _ := store.Get(r, "CUSupermileage")

	// Get authenticated session value from cookie
	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		// Require login
		loginHandler(w, r)
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

// apiHandler provides access to data stored in the DB
func apiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" && r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	// Response message string
	resp := ""

	switch r.URL.Path {
	case "/carName":
		resp = api.CarNameHandler(r)
		break

	case "/csv":
		resp = api.CSVHandler(r)
		break

	case "/graph":
		resp = api.GraphHandler(r)
		break

	case "/latestRunRow":
		resp = api.LatestRunHandler(r)
		break

	case "/runIds":
		resp = api.RunIdsHandler(r)
		break
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

// loginHandler handles requests to the login page
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed "+r.Method, http.StatusMethodNotAllowed)
		return
	}

	// Values that will be used in the template
	vals := map[string]string{}

	// Parse the root page template
	t, _ := template.ParseFiles("templates/login.html")

	// Respond with the template filled with the values
	t.Execute(w, vals)
}

// authHandler handles authentication
func authHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get storage cookie store
	session, _ := store.Get(r, "CUSupermileage")

	// Get authenticated session value from cookie
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		// Already authorized, no need to do anything further
		return
	}

	// Get values from request
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	// Simple password validation
	if username != "cedarville" || password != "bebold" {
		http.Error(w, "Invalid username or password!", http.StatusBadRequest)
		return
	}

	// Passed validation, set session
	session.Values["authenticated"] = true
	session.Save(r, w)

	log.Info("Logged in!")

	http.Redirect(w, r, "/", http.StatusFound)
}
