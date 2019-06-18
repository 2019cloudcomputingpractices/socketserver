package httpparser

import (
	"fmt"
	"strings"
)

type RequestMessage struct {
	requestline Requestline
	header      Header
	requestbody Requestbody
}

func ParseRequestMessage(mess_str string) (RequestMessage, error) {
	var RequestMessage RequestMessage
	lines := strings.Split(mess_str, "\n")
	//parse request line
	var err error
	RequestMessage.requestline, err = ParseRequestLine(lines[0])
	if err != nil {
		return RequestMessage, err
	}
	index := 1
	for ; index < len(lines) && strings.Trim(lines[index], " \n\t\r") != ""; index++ {
		continue
	}
	RequestMessage.header, err = ParseHeader(lines[1:index])
	for ; index < len(lines) && strings.Trim(lines[index], " \n\t\r") == ""; index++ {
		continue
	}
	RequestMessage.requestbody = Requestbody(strings.Join(lines[index:], "\n"))
	return RequestMessage, err
}

func (m RequestMessage) ToString() string {
	return fmt.Sprintf("%s\n%s\n%s\n", m.requestline.ToString(), m.header.ToString(), string(m.requestbody))
}

func (m RequestMessage) GetMethod() string {
	return m.requestline.method
}

func (m RequestMessage) GetUrl() string {
	return m.requestline.url
}

func (m RequestMessage) GetBody() Requestbody {
	return m.requestbody
}
