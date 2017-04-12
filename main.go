// Package main is all that's needed to run my site. It's little more
// than a simple http server that deals with one URL (but it does it well!).
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	// Static files
	app.Static("/img", "./img")
	app.Static("/fonts", "./fonts")
	app.Static("/js", "./js")
	app.Static("/css", "./css")
	app.Static("/dice", "./dice/app")
	app.Static("/static", "./static")

	// HTML + Templates
	app.LoadHTMLGlob("html/*")

	// Routes
	app.GET("/ping", pinger)
	app.GET("/quotes", quoter)
	app.POST("/dice/roll", roller)
	app.GET("/", index)

	// Engage
	app.Run(":3000")
}
