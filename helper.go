package gologs

import (
	"fmt"
	"log"
)

func logPrefix(l *Logger, t string) string {
	var clr string

	switch t {
	case "info":
		clr = l.Colors.Info
	case "warn":
		clr = l.Colors.Warning
	case "err":
		clr = l.Colors.Error
	case "ftl":
		clr = l.Colors.Fatal
	default:
		log.Fatalln("Error en el tipo ingresado")
	}

	return fmt.Sprintf("%s (v: %s) - [%s] - ", l.apiName, l.apiVersion, getColor(clr, t))
}

func getColor(c string, i interface{}) string {
	return fmt.Sprintf("%s%s%s", c, i, reset)
}
