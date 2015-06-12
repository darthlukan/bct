package main

import (
	"github.com/darthlukan/bct/quotes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}

func pinger(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "PONG!"})
}

func quoter(context *gin.Context) {
	quote := quotes.RandomQuote()
	if len(quote) > 0 {
		context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": quote})
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "data": "Error, quote not found."})
	}
}
