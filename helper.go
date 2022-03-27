package gologs

import (
	"fmt"
	"github.com/rmacciuci/gologs/colors"
)

func logPrefix(l *Logger, t string) string {
	var clr string

	switch t {
	case "info":
		clr = l.colors.Info
	case "warn":
		clr = l.colors.Warning
	case "err":
		clr = l.colors.Error
	case "ftl":
		clr = l.colors.Fatal
	default:
		clr = l.colors.Info
	}

	return fmt.Sprintf("%s (v: %s) - [%s] - ", l.apiName, l.apiVersion, getColor(clr, t))
}

func getColor(c string, i interface{}) string {
	return fmt.Sprintf("%s%s%s", c, i, colors.GetReset())
}
