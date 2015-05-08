// Package main is all that's needed to run my site. It's little more
// than a simple http server that deals with one URL (but it does it well!).
package main

import (
	"github.com/darthlukan/bct/quotes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	app := gin.Default()

	// Static files
	app.Static("/img", "./img")
	app.Static("/fonts", "./fonts")
	app.Static("/js", "./js")
	app.Static("/css", "./css")

	// HTML + Templates
	app.LoadHTMLGlob("html/*")

	// Routes
	app.GET("/ping", pinger)
	app.GET("/quotes", quoter)
	app.GET("/", index)

	// Engage
	app.Run(":3000")
}

func index(context *gin.Context) {
	index := "index.html"
	context.HTML(http.StatusOK, index, gin.H{})
}

func pinger(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "OK!", "data": "PONG!"})
}

func quoter(context *gin.Context) {
	quote := quotes.RandomQuote()
	if len(quote) > 0 {
		context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": quote})
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "Failed", "data": "Error, quote not found."})
	}
}
