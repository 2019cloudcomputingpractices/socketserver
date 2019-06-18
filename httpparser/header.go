package httpparser

import (
	"fmt"
	"strings"
)

type Header map[string]string

func ParseHeader(header []string) (Header, error) {
	headerMap := make(map[string]string)
	for _, h := range header {
		ele := strings.Split(h, ":")
		key := strings.Trim(ele[0], " \r\n\t")
		value := strings.Trim(strings.Join(ele[1:], ":"), " \r\n\t")
		//fmt.Printf("key : %s, value: %s\n", key, value)
		headerMap[key] = value
	}
	return Header(headerMap), nil
}

func (h Header) ToString() string {
	str := ""
	for k, v := range h {
		str += fmt.Sprintf("%s : %s\n", k, v)
	}
	return str
}

func (h Header) SetHeader(key string, value string) {
	h[key] = value
}
