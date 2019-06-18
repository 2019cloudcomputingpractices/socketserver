package httpparser

import (
	"fmt"
)

type Responseline struct {
	version     string
	status      int
	description string
}

func (l Responseline) ToString() string {
	return fmt.Sprintf("%s %d %s", l.version, l.status, l.description)
}

func GetResponseLine(version string, status int, description string) Responseline {
	return Responseline{
		version:     version,
		status:      status,
		description: description,
	}
}
