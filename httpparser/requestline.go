package httpparser

import (
	"errors"
	"fmt"
	"strings"
)

type Requestline struct {
	method  string
	url     string
	version string
}

func ParseRequestLine(line_str string) (Requestline, error) {
	var requestline Requestline
	element := strings.Split(line_str, " ")
	if len(element) != 3 {
		return requestline, errors.New("invalid request line format")
	}
	requestline.method = strings.Trim(element[0], " \r\n\t")
	requestline.url = strings.Trim(element[1], " \r\n\t")
	requestline.version = strings.Trim(element[2], " \r\n\t")
	return requestline, nil
}

func (line Requestline) ToString() string {
	return fmt.Sprintf("%s %s %s", line.method, line.url, line.version)
}
