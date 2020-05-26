package simplelog

import "log"

type Log struct{}

func (s *Log) Init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
}

func (s *Log) Debug(v string) {
	log.Println(v)
}

func (s *Log) Info(v string) {
	log.Println(v)
}

func (s *Log) Error(err error) {
	log.Println(err.Error())
}

func (s *Log) Warning(v string) {
	log.Println(v)
}

func (s *Log) Debugf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (s *Log) Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (s *Log) Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (s *Log) Warningf(format string, args ...interface{}) {
	log.Printf(format, args...)
}
