package rmlogger

var Logs Logger

type Logger struct {
	colors     *colors
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

	return &Logs
}
