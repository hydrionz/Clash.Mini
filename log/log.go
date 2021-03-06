package log

import (
	cLog "github.com/Dreamacro/clash/log"
)

func Infoln(format string, v ...interface{}) {
	cLog.Infoln(format, v...)
}

func Warnln(format string, v ...interface{}) {
	cLog.Warnln(format, v...)
}

func Errorln(format string, v ...interface{}) {
	cLog.Errorln(format, v...)
}

func Debugln(format string, v ...interface{}) {
	cLog.Debugln(format, v...)
}

func Fatalln(format string, v ...interface{}) {
	cLog.Fatalln(format, v...)
}

func Level() cLog.LogLevel {
	return cLog.Level()
}

func SetLevel(newLevel cLog.LogLevel) {
	cLog.SetLevel(newLevel)
}
