package utils

import (
	"log"
	"os"
)

type Singgerx struct {
	logger *log.Logger
}

var Singger Singgerx

func InitLogger() {
	Singger.logger = log.New(os.Stdout, "[INFO]:", log.Lshortfile)
}

func (s *Singgerx) Debug(str string) {
	s.logger.SetPrefix("[DEBUG]:")
	s.logger.Println(str)
}

func (s *Singgerx) Info(str string) {
	s.logger.SetPrefix("[INFO]:")
	s.logger.Println(str)

}

func (s *Singgerx) Fatal(str string) {
	s.logger.SetPrefix("[FATAL]:")
	s.logger.Println(str)

}
