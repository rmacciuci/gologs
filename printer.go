package rmlogger

import "log"

func (l *Logger) Info(d interface{}) {
	log.SetPrefix(logPrefix(l, "info"))
	log.Println(d)
}

func (l *Logger) Warning(d interface{}) {
	log.SetPrefix(logPrefix(l, "warn"))
	log.Println(d)
}

func (l *Logger) Error(d interface{}) {
	log.SetPrefix(logPrefix(l, "err"))
	log.Println(d)
}

func (l *Logger) Fatal(d interface{}) {
	log.SetPrefix(logPrefix(l, "ftl"))
	log.Fatalln(d)
}
