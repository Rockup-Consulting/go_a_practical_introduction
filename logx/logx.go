package logx

import (
	"fmt"
	"log"
	"os"
)

func New(serviceName string, verbose bool) *log.Logger {
	file, err := os.Create("tmp.logs")
	if err != nil {
		fmt.Println(err)
	}
	l := log.New(file, serviceName+": ", log.Ldate|log.Ltime|log.LUTC|log.Lmsgprefix)

	if verbose {
		l.SetOutput(os.Stdout)
	}

	return l
}
