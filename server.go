/*
#!/usr/bin/env gorun
@author :yinzhengjie
Blog:http://www.cnblogs.com/yinzhengjie/tag/GO%E8%AF%AD%E8%A8%80%E7%9A%84%E8%BF%9B%E9%98%B6%E4%B9%8B%E8%B7%AF/
EMAIL:y1053419035@qq.com
*/

package main

import (
	"fmt"
	"log"
	"net"
	"socketserver/httpparser"
	"socketserver/router"
)

var content = `HTTP/1.1 200 OK
Date: Sat, 29 Jul 2017 06:18:23 GMT
Content-Type: text/html
Connection: Keep-Alive
Server: BWS/1.1
X-UA-Compatible: IE=Edge,chrome=1
BDPAGETYPE: 3
Set-Cookie: BDSVRTM=0; path=/

test
`

func Handle_conn(conn net.Conn) { //这个是在处理客户端会阻塞。
	//fmt.Println("handle one connection")
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	//fmt.Println("get result success")
	if err != nil {
		fmt.Println(err.Error())
		conn.Close()
		return
	}
	//fmt.Println(string(buf))
	mess, err := httpparser.ParseRequestMessage(string(buf))
	if err != nil {
		fmt.Println(err.Error())
		conn.Close()
		return
	}
	fmt.Println(string(mess.GetBody()))
	content = router.HandleMessage(&mess)
	//fmt.Println("the return content is:")
	//fmt.Println(content)
	conn.Write([]byte(content)) //将html的代码返回给客户端，这样客户端在web上访问就可以拿到指定字符。
	conn.Close()
}

func main() {
	addr := "0.0.0.0:7777" //表示监听本地所有ip的8080端口，也可以这样写：addr := ":8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept() //用conn接收链接
		if err != nil {
			log.Fatal(err)
		}
		go Handle_conn(conn) //将接受来的链接交给该函数去处理。
	}
}
