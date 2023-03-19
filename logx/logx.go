package logx

import (
	"fmt"
	"log"
	"os"
)

const logFlags = log.Ldate | log.Ltime | log.LUTC | log.Lmsgprefix

func New(serviceName string, verbose bool) *log.Logger {
	if verbose {
		return log.New(os.Stdout, serviceName+": ", logFlags)
	} else {
		file, err := os.Create("tmp.logs")
		if err != nil {
			fmt.Println(err)
		}
		return log.New(file, serviceName+": ", logFlags)
	}
}
