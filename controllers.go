package main

import (
	"github.com/darthlukan/bct/quotes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func index(context *gin.Context) {
	index := "index.html"
	context.HTML(http.StatusOK, index, gin.H{})
}

func pinger(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "PONG!"})
	return
}

func quoter(context *gin.Context) {
	quote := quotes.RandomQuote()
	if len(quote) > 0 {
		context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": quote})
		return
	}
	context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "Error, quote not found."})
	return
}

type Dice struct {
	D4  int `json:"d4"`
	D6  int `json:"d6"`
	D8  int `json:"d8"`
	D10 int `json:"d10"`
	D12 int `json:"d12"`
	D20 int `json:"d20"`
}

func roller(context *gin.Context) {
	var dice Dice

	context.Bind(&dice)
}
