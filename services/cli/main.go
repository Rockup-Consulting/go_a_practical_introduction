package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"todo/logx"
	"todo/todo"
)

func main() {
	verbose := flag.Bool("verbose", false, "pass the verbose flag to print logs to stdout")
	flag.Parse()

	l := logx.New("TodoList", *verbose)

	if err := run(l); err != nil {
		fmt.Println("Unexpected error, quitting TodoList")
		l.Printf("ERROR: program exiting: %s", err)
		os.Exit(1)
	}
}

const startupInstruction string = `TodoList: The ultimate check list utility

Usage:
l        - list all Todo Items
lc       - list all Checked Todo Items
lu       - list all Unchecked Todo Items
t [ID]   - toggle Todo Item with ID [ID]
d [ID]   - delete Todo Item with ID [ID]
c [Msg]  - create new Todo Item with message [Msg]
q        - quit TodoList
h        - display this menu again
--------------------------------------------------
`

func run(l *log.Logger) error {

	fmt.Print(startupInstruction)

	store, err := initFileStore(l)
	if err != nil {
		return err
	}

	defer func() {
		l.Println("start closing store file")
		if closeErr := store.Close(); closeErr != nil {
			l.Printf("ERROR: closing file: %s", err)
		} else {
			l.Println("store file close success")
		}
	}()

	l.Println("creating TodoStore")
	todoStore := todo.NewFsStore(store, l)
	routes := InitRoutes(todoStore, l)

	l.Println("creating new buffered input reader")
	inputReader := bufio.NewReader(os.Stdin)

	l.Println("listening for standard input")
	for {
		in, err := inputReader.ReadString('\n')
		if err != nil {
			return err
		}

		if in[len(in)-1] == '\n' {
			in = in[:len(in)-1]
		}

		instruction, arg, ok := strings.Cut(in, " ")
		if ok {
			if err := routes.HandleInstructionWithArg(instruction, arg); err != nil {
				return err
			}
			continue
		}

		if instruction == "q" {
			fmt.Println("see you soon!")
			break
		}

		if err := routes.HandleInstruction(instruction); err != nil {
			return err
		}
	}

	return nil
}

func initFileStore(l *log.Logger) (*os.File, error) {
	storeLocation := os.Getenv("TODO_LIST_STORE_LOCATION")
	if storeLocation == "" {
		storeLocation = "tmp.store.json"
	}

	l.Printf("store location is %q", storeLocation)

	l.Printf("opening file store")
	store, err := os.OpenFile(storeLocation, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	// ensure that store is not an empty file, an empty file is not valid JSON
	if _, err := store.Read(make([]byte, 1)); err != nil {
		if err == io.EOF {
			l.Printf("bootstrapping empty JSON file")
			// write an empty list to the JSON file
			store.Write([]byte("[]"))

			// reset the file cursor to the start
			store.Seek(0, 0)
		} else {
			return nil, err
		}
	}

	return store, nil
}
