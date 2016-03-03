package logger

import (
	"log"
	"os"
)

var logger2 *log.Logger

func init() {
	logger2 = log.New(os.Stdout, "[INFO]:", log.Lshortfile)
}

func Debug(str string) {
	logger2.SetPrefix("[DEBUG]:")
	logger2.Println(str)
}

func Info(str string) {
	logger2.SetPrefix("[INFO]:")
	logger2.Println(str)
}

func Fatal(str string) {
	logger2.SetPrefix("[FATAL]:")
	logger2.Println(str)
}
