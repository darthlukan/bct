package controllers

import (
	"github.com/laurent22/ripple"
)

type PingController struct {
	Pong string
}

func NewPingController() *PingController {
	ping := new(PingController)
	ping.Pong = "Pong!"
	return ping
}

func (self *PingController) Get(ctx *ripple.Context) {
	ctx.Response.Body = self.Pong
}
