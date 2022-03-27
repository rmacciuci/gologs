package gologs

import (
	"fmt"
	"github.com/rmacciuci/gologs/pkg/colors"
)

func logPrefix(l *Logger, t string) string {

	return fmt.Sprintf("%s (v: %s) - [%s] - ", l.apiName, l.apiVersion, l.getTagColor(t))
}

func (l *Logger) getTagColor(t interface{}) string {
	var c string

	switch t {
	case "info":
		c = l.colors.Info
	case "warn":
		c = l.colors.Warning
	case "error":
		c = l.colors.Error
	case "fatal":
		c = l.colors.Fatal
	default:
		c = l.colors.Info
	}

	return fmt.Sprintf("%s%s%s", c, t, colors.GetReset())
}
