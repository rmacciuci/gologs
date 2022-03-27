package gologs

import (
	"github.com/Rmacciuci/gologs/colors"
)

type Logger struct {
	colors     *colors.Colors
	apiName    string
	apiVersion string
}

func Define(an string, av string) Logger {
	l := Logger{}

	if an != "" {
		l.apiName = an
	}

	if av != "" {
		l.apiVersion = av
	}

	// Set Colors
	l.colors = colors.InitDefault()

	return l
}
