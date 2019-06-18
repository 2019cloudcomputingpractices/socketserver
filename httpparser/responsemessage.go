package httpparser

import (
	"fmt"
)

type ResponseMessage struct {
	responseline Responseline
	header       Header
	responsebody Responsebody
}

func (m ResponseMessage) ToString() string {
	return fmt.Sprintf("%s\n%s\n%s", m.responseline.ToString(), m.header.ToString(), m.responsebody)
}

func GetResponseMessage(responseline Responseline, header Header, responsebody Responsebody) ResponseMessage {
	return ResponseMessage{
		responseline: responseline,
		header:       header,
		responsebody: responsebody,
	}
}
