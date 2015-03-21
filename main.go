// Package main is all that's needed to run my site. It's little more
// than a simple http server that deals with one URL (but it does it well!).
package main

import (
	"bufio"
	"fmt"
	"github.com/Radiobox/web_responders"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/objx"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

var (
	projectRoot string
	templates   *template.Template
	goPath      = os.Getenv("GOPATH")
	messages    = web_responders.NewMessageMap()
	content     = objx.Map{
		"home": objx.Map{
			"one":   "",
			"two":   "",
			"three": "",
			"four":  "",
			"big":   "",
		},
	}
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
	goweb.MapStatic("/static_content", path.Join(projectRoot, "static_content"))

	address := ":3000"
	if port := os.Getenv("PORT"); port != "" {
		address = ":" + port
	}
	server := &http.Server{
		Addr:           address,
		Handler:        &LoggedHandler{goweb.DefaultHttpHandler()},
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

type LoggedHandler struct {
	baseHandler http.Handler
}

func (handler *LoggedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	go log.Printf("%s Request for %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
	handler.baseHandler.ServeHTTP(w, r)
}

// TODO: Do I really want this much I/O per request?
func readFile(path string) string {
	var staticBase string = "static_content"
	path = fmt.Sprintf("%v/%v/%v", projectRoot, staticBase, path)
	file, err := os.Open(path)

	if err != nil {
		return err.Error()
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var text string
	text = strings.Join(lines, "\n")

	return text
}

func handleTemplate(ctx context.Context, template string) error {
	// TODO: Populate "content"
	if err := templates.ExecuteTemplate(ctx.HttpResponseWriter(), template, content); err != nil {
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
