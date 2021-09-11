package log

import "fmt"

type Logger interface {
	Log(mess string)
	Notify(mess string) bool
	Start() error
}

type DefaultLogger struct{}

func (logger *DefaultLogger) Log(mess string) {
	fmt.Println(mess)
}

func (logger *DefaultLogger) Notify(mess string) bool {
	_, err := fmt.Println(mess)
	return err == nil
}

func (logger *DefaultLogger) Start() error {
	return nil
}
