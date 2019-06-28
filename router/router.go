package router

import (
	"github.com/2019cloudcomputingpractices/socketserver/controller"
	"github.com/2019cloudcomputingpractices/socketserver/httpparser"
)

func HandleMessage(mess *httpparser.RequestMessage) string {
	if mess.GetUrl() == "/register" {
		return controller.HandleRegister(mess)
	} else if mess.GetUrl() == "/login" {
		return controller.HandleLogin(mess)
	} else {
		return controller.HandleStatic(mess)
	}
}
