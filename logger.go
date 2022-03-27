package gologs

import (
	"github.com/rmacciuci/gologs/pkg/colors"
)

type printerSettings struct {
	infoPrint bool
	warnPrint bool
	errPrint  bool
	ftlPrint  bool
}

type Logger struct {
	printerSettings
	colors     *colors.Colors
	apiName    string
	apiVersion string
}

func Define(an string, av string) Logger {
	l := Logger{
		printerSettings: printerSettings{true, true, true, true},
	}

	if an != "" {
		l.apiName = an
	}

	if av != "" {
		l.apiVersion = av
	}

	// Set Colors
	l.colors = colors.InitDefault()

	welcomeMessage(l)

	return l
}

func welcomeMessage(l Logger) {
	l.Info("Init GOLOGS By Ramiro Macciuci")
}
