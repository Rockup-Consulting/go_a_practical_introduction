package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"todo/logx"
)

func main() {
	verbose := flag.Bool("verbose", false, "pass the verbose flag to print logs to stdout")
	flag.Parse()

	l := logx.New("TodoList", *verbose)

	if err := run(l); err != nil {
		l.Printf("ERROR: program exiting: %s", err)
		os.Exit(1)
	}
}

func run(l *log.Logger) error {
	return errors.New("something really went wrong")
}
