// Package main is all that's needed to run my site. It's little more
// than a simple http server that deals with one URL (but it does it well!).
package main

import (
	//"github.com/gorilla/handlers"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"github.com/Radiobox/web_responders"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"time"
)

type LoggedHandler struct {
	baseHandler http.Handler
}

func (handler *LoggedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	go log.Printf("%s Request for %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
	handler.baseHandler.ServeHTTP(w, r)
}

func handleTemplate(ctx context.Context, template string) error {
	if err := templates.ExecuteTemplate(ctx.HttpResponseWriter(), template, ctx.HttpRequest()); err != nil {
		messages.AddErrorMessage("Could not load template " + template + ": " + err.Error())
		return goweb.Respond.With(ctx, http.StatusInternalServerError, []byte(err.Error()))
	}
	return nil
}

func htmlFileHandler(ctx context.Context) error {
	template := path.Join(ctx.Path().Segments()[1:]...) + ctx.FileExtension()
	return handleTemplate(ctx, template)
}

// IndexHandler checks if the URL.Path contains a dot (.) as in the case
// of filetypes in the URL.  In this case, it uses a different Mux
// to handle such requests, else it uses the default Mux and executes
// the index template.
func indexHandler(ctx context.Context) error {
	return handleTemplate(ctx, "index.html")
}

var (
	projectRoot string
	templates   *template.Template
	goPath      = os.Getenv("GOPATH")
	messages    = web_responders.NewMessageMap()
	//myMux       = http.NewServeMux()
)

func main() {
	log.Println("Starting server...")

	if goPath == "" {
		projectRoot = "."
	} else {
		projectRoot = path.Join(goPath, "src", "github.com", "darthlukan", "bct")
	}

	templates = template.Must(template.ParseGlob(projectRoot + "/html/*"))
	goweb.Map("/", indexHandler)
	goweb.Map("/html/***", htmlFileHandler)
	goweb.MapStatic("/css", path.Join(projectRoot, "css"))
	goweb.MapStatic("/js", path.Join(projectRoot, "js"))
	goweb.MapStatic("/img", path.Join(projectRoot, "img"))

	//myMux.Handle("/", http.FileServer(http.Dir(projectRoot)))

	//http.HandleFunc("/", indexHandler)

	address := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		address = ":" + port
	}
	server := &http.Server{
		Addr: address,
		Handler: &LoggedHandler{goweb.DefaultHttpHandler()},
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	listener, listenErr := net.Listen("tcp", address)
	if listenErr != nil {
		log.Panicf("Could not listen for TCP on %s: %s", address, listenErr)
	}
	log.Println("Server loaded, check localhost" + address)
	server.Serve(listener)
}
