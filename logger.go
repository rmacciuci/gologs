package rmlogger

type Logger struct {
	colors     *colors
	apiName    string
	apiVersion string
}

func Define(an string, av string) *Logger {
	logs := Logger{}

	if an != "" {
		logs.apiName = an
	}

	if av != "" {
		logs.apiVersion = av
	}

	return &logs
}
