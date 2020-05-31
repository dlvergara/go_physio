package logger

import (
	"physiobot/config"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	success *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	fatal   *log.Logger
)

func init() {
	configuration := config.GetConfig()
	out := ioutil.Discard
	eout := ioutil.Discard

	if configuration.LogEnabled {
		out = os.Stdout
		eout = os.Stderr
	}

	success = log.New(out,
		"\033[32m\n[OK]: \033[39m",
		log.Ldate|log.Ltime|log.Lshortfile)

	info = log.New(out,
		"\033[36m\n[INFO]: \033[39m",
		log.Ldate|log.Ltime|log.Lshortfile)

	warning = log.New(out,
		"\033[33m\n[WARNING]: \033[39m",
		log.Ldate|log.Ltime|log.Lshortfile)

	err = log.New(eout,
		"\033[31m\n[ERROR]: \033[39m",
		log.Ldate|log.Ltime|log.Lshortfile)

	fatal = log.New(eout,
		"\033[31m\n[FATAL]: \033[39m",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func format(message interface{}) string {
	return "\n" + fmt.Sprintf("%+v", message) + "\n\n"
}

// Error log
func Error(message interface{}) {
	err.Output(2, format(message))
}

// Info log
func Info(message interface{}) {
	info.Output(2, format(message))
}

// Warning log
func Warning(message interface{}) {
	warning.Output(2, format(message))
}

// Success log
func Success(message interface{}) {
	success.Output(2, format(message))
}

// Fatal log
func Fatal(message interface{}) {
	fatal.Output(2, format(message))
	os.Exit(1)
}
