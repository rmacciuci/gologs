package gologs

import (
	"log"
)

func (l *Logger) Info(d interface{}) {
	if !l.infoPrint {
		return
	}
	log.SetPrefix(logPrefix(l, "info"))
	log.Println(d)
}

func (l *Logger) Warning(d interface{}) {
	if !l.warnPrint {
		return
	}
	log.SetPrefix(logPrefix(l, "warn"))
	log.Println(d)
}

func (l *Logger) Error(d interface{}) {
	if !l.errPrint {
		return
	}
	log.SetPrefix(logPrefix(l, "error"))
	log.Println(d)
}

func (l *Logger) Fatal(d interface{}) {
	if !l.ftlPrint {
		return
	}
	log.SetPrefix(logPrefix(l, "fatal"))
	log.Fatalln(d)
}

func (l *Logger) InfoStatus(s bool) {
	l.infoPrint = s
}

func (l *Logger) WarnStatus(s bool) {
	l.warnPrint = s
}

func (l *Logger) ErrorStatus(s bool) {
	l.errPrint = s
}

func (l *Logger) FatalStatus(s bool) {
	l.ftlPrint = s
}
