// Package main is all that's needed to run my site. It's little more
// than a simple http server that deals with one URL (but it does it well!).
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	app := gin.Default()
	app.GET("/ping", pinger)
	app.Run(":3000")
}

func pinger(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "OK!", "data": "PONG!"})
}
