package httpparser

import (
	"strings"
)

type Requestbody string

func ParseForm(body Requestbody) map[string]string {
	formMap := make(map[string]string)
	str := strings.Trim(string(body), " \n\r\t\x00")
	statments := strings.Split(str, "&")
	for _, e := range statments {
		ele := strings.Split(e, "=")
		//fmt.Println(ele)
		//fmt.Printf("%s : %s\n", ele[0], ele[1])
		formMap[ele[0]] = ele[1]
	}
	return formMap
}
