// Package main is all that's needed to run my site. It's little more
// than a simple http server that deals with one URL (but it does it well!).
package main

import (
	"github.com/darthlukan/bct/controllers"
	"github.com/laurent22/ripple"
	"log"
	"net/http"
	"strings"
)

func main() {
	app := ripple.NewApplication()
	pingController := controllers.NewPingController()
	app.RegisterController("ping", pingController)

	// TODO: Properly split out the site handler from the API/Resource handler
	//app.AddRoute(ripple.Route{Pattern: "/"})
	app.AddRoute(ripple.Route{Pattern: ":_controller"})
	//http.HandleFunc("/", TemplatedHandler)
	app.SetBaseUrl("/")
	http.HandleFunc("/", app.ServeHTTP)

	log.Printf("Server loaded on localhost:3000\n")
	http.ListenAndServe(":3000", &LoggedHttpHandler{app})
}

type LoggedHttpHandler struct {
	baseHandler http.Handler
}

func (handler *LoggedHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	go log.Printf("%s Request for %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
	handler.baseHandler.ServeHTTP(w, r)
}

func TemplatedHandler(w http.ResponseWriter, r *http.Request) {
	var content []byte

	splitPath := strings.Split(r.URL.Path, "/")
	htmlName := splitPath[len(splitPath)-1]
	if htmlName == "" {
		content = []byte("Put a pretty HTML index here")
	} else {
		content = []byte(htmlName) // Load named HTML
	}
	w.Write(content)
}
