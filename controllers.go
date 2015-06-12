package main

import (
	"github.com/darthlukan/bct/quotes"
	"github.com/darthlukan/dice/roll"
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
	context.JSON(http.StatusNotFound, gin.H{
		"status": http.StatusNotFound,
		"data":   "Error, quote not found."})
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
	var die Dice

	context.Bind(&die)
	d4 := roll.Roll(int64(4), int64(die.D4))
	d6 := roll.Roll(int64(6), int64(die.D6))
	d8 := roll.Roll(int64(8), int64(die.D8))
	d10 := roll.Roll(int64(10), int64(die.D10))
	d12 := roll.Roll(int64(12), int64(die.D12))
	d20 := roll.Roll(int64(20), int64(die.D20))
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": map[string]int64{
			"d4":  d4,
			"d6":  d6,
			"d8":  d8,
			"d10": d10,
			"d12": d12,
			"d20": d20,
		},
	})
	return
}
