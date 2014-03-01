// Package main is all that's needed to run my site. It's little more
// than a simple http server that deals with one URL (but it does it well!).
package main

import (
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"time"
)

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

func indexHandler(ctx context.Context) error {
	return handleTemplate(ctx, "index.html")
}

var (
	projectRoot string
	templates   *template.Template
	goPath      = os.Getenv("GOPATH")
	messages    = web_responders.NewMessageMap()
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

	address := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		address = ":" + port
	}
	server := &http.Server{
		Addr:           address,
		Handler:        goweb.DefaultHttpHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	listener, listenErr := net.Listen("tcp", address)
	if listenErr != nil {
		log.Panicf("Could not listen for TCP on %s: %s", address, listenErr)
	}
	log.Println("Server loaded, check localhost" + address)
	server.Serve(listener)
}
