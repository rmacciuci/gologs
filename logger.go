package gologs

var Logs Logger

type Logger struct {
	Colors     *colors
	apiName    string
	apiVersion string
}

func Define(an string, av string) *Logger {
	Logs = Logger{}

	if an != "" {
		Logs.apiName = an
	}

	if av != "" {
		Logs.apiVersion = av
	}

	// Set Colors
	Logs.Colors = NewColors()

	return &Logs
}
