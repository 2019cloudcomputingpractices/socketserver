package controller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"text/template"

	"github.com/2019cloudcomputingpractices/socketserver/httpparser"
	"github.com/2019cloudcomputingpractices/socketserver/model"
)

func HandleRegister(mess *httpparser.RequestMessage) string {
	if mess.GetMethod() == "GET" || mess.GetMethod() == "HEAD" {
		content, err := ioutil.ReadFile("./static/register.html")
		if err != nil {
			fmt.Println("file not found")
			return NotFound()
		}
		responseline := httpparser.GetResponseLine("HTTP/1.1", 200, "OK")
		header := make(map[string]string)
		header["Content-Type"] = "text/html; charset=utf-8"
		responsebody := httpparser.Responsebody("")
		if mess.GetMethod() == "GET" {
			responsebody = httpparser.Responsebody(content)
		}
		mess := httpparser.GetResponseMessage(responseline, httpparser.Header(header), responsebody)
		return mess.ToString()
	} else if mess.GetMethod() == "POST" {
		formMap := httpparser.ParseForm(mess.GetBody())
		username := formMap["username"]
		password := formMap["password"]
		fmt.Println("username: "+username+", password: ", password)
		err := model.CreateUser(username, password)
		info := "注册成功"
		if err != nil {
			info = err.Error()
		}
		responseline := httpparser.GetResponseLine("HTTP/1.1", 200, "OK")
		header := make(map[string]string)
		header["Content-Type"] = "text/html; charset=utf-8"
		t, _ := template.ParseFiles("static/info.html")
		buf := bytes.NewBufferString("")
		t.Execute(buf, info)
		responsebody := httpparser.Responsebody(buf.String())
		mess := httpparser.GetResponseMessage(responseline, httpparser.Header(header), responsebody)
		return mess.ToString()
	} else {
		return NotFound()
	}
}

func HandleLogin(mess *httpparser.RequestMessage) string {
	if mess.GetMethod() == "GET" || mess.GetMethod() == "HEAD" {
		content, err := ioutil.ReadFile("./static/login.html")
		if err != nil {
			fmt.Println("file not found")
			return NotFound()
		}
		responseline := httpparser.GetResponseLine("HTTP/1.1", 200, "OK")
		header := make(map[string]string)
		header["Content-Type"] = "text/html; charset=utf-8"
		responsebody := httpparser.Responsebody("")
		if mess.GetMethod() == "GET" {
			responsebody = httpparser.Responsebody(content)
		}
		mess := httpparser.GetResponseMessage(responseline, httpparser.Header(header), responsebody)
		return mess.ToString()
	} else if mess.GetMethod() == "POST" {
		fmt.Println("controller handle post user")
		formMap := httpparser.ParseForm(mess.GetBody())
		username := formMap["username"]
		password := formMap["password"]
		fmt.Println("username: "+username+", password: ", password)
		err := model.Login(username, password)
		info := "登录成功"
		if err != nil {
			info = err.Error()
		}
		responseline := httpparser.GetResponseLine("HTTP/1.1", 200, "OK")
		header := make(map[string]string)
		header["Content-Type"] = "text/html; charset=utf-8"
		t, _ := template.ParseFiles("static/info.html")
		buf := bytes.NewBufferString("")
		t.Execute(buf, info)
		responsebody := httpparser.Responsebody(buf.String())
		mess := httpparser.GetResponseMessage(responseline, httpparser.Header(header), responsebody)
		return mess.ToString()
	} else {
		return NotFound()
	}
}

func HandleStatic(mess *httpparser.RequestMessage) string {
	if mess.GetMethod() == "GET" || mess.GetMethod() == "HEAD" {
		content, err := ioutil.ReadFile("./static" + mess.GetUrl())
		if err != nil {
			fmt.Println("file not found")
			return NotFound()
		}
		responseline := httpparser.GetResponseLine("HTTP/1.1", 200, "OK")
		header := make(map[string]string)
		header["Content-Type"] = "text/html; charset=utf-8"
		responsebody := httpparser.Responsebody("")
		if mess.GetMethod() == "GET" {
			responsebody = httpparser.Responsebody(content)
		}
		mess := httpparser.GetResponseMessage(responseline, httpparser.Header(header), responsebody)
		return mess.ToString()
	} else {
		return NotFound()
	}
}

func NotFound() string {
	responseline := httpparser.GetResponseLine("HTTP/1.1", 404, "Not Found")
	header := make(map[string]string)
	header["Content-Type"] = "text/html; charset=utf-8"
	responsebody := httpparser.Responsebody("404 not found")
	mess := httpparser.GetResponseMessage(responseline, httpparser.Header(header), responsebody)
	return mess.ToString()
}
