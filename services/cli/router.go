package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"todo/todo"
)

type Routes struct {
	todoStore todo.Store
	l         *log.Logger
	tw        *tabwriter.Writer
}

func InitRoutes(todoStore todo.Store, l *log.Logger) Routes {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	return Routes{todoStore, l, tw}
}

func (r Routes) HandleInstruction(instruction string) error {
	switch instruction {
	case "l":
		r.l.Println("list all Todo Items")
		list, err := r.todoStore.List(todo.All)
		if err != nil {
			return err
		}

		r.writeTodoList(list)
	case "lc":
		r.l.Println("list all Checked Todo Items")

		list, err := r.todoStore.List(todo.Checked)
		if err != nil {
			return err
		}

		r.writeTodoList(list)
	case "lu":
		r.l.Println("list all Unchecked Todo Items")

		list, err := r.todoStore.List(todo.NotChecked)
		if err != nil {
			return err
		}

		r.writeTodoList(list)
	case "t":
		r.l.Println("invalid input: toggle command expects exactly one argument")
		fmt.Println("toggle Todo Item requires exactly one argument: TodoItem ID")
	case "d":
		r.l.Println("invalid input: delete command expects exactly one argument")
		fmt.Println("delete Todo Item requires exactly one argument: TodoItem ID")
	case "c":
		r.l.Println("invalid input: create command expects exactly one argument")
		fmt.Println("create new Todo Item requires exactly one argument: TodoItem Message")
	case "h":
		fmt.Print(startupInstruction)
	default:
		r.l.Printf("unknown command: %q\n", instruction)
		fmt.Printf("unknown command: %q\n", instruction)
	}

	return nil
}

func (r Routes) HandleInstructionWithArg(instruction, arg string) error {
	switch instruction {
	case "l":
		r.l.Println("invalid input: list command expects zero arguments")
		fmt.Println("list command expects zero arguments")

	case "lc":
		r.l.Println("invalid input: list checked command expects zero arguments")
		fmt.Println("list checked command expects zero arguments")

	case "lu":
		r.l.Println("invalid input: list unchecked command expects zero arguments")
		fmt.Println("list unchecked command expects zero arguments")

	case "t":
		err := r.todoStore.Toggle(arg)
		if err != nil {
			if todo.IsNotFoundErr(err) {
				fmt.Printf("TodoItem %q not found", arg)
				return nil
			}

			return err
		}

	case "d":
		err := r.todoStore.Delete(arg)
		if err != nil {
			if todo.IsNotFoundErr(err) {
				fmt.Printf("TodoItem %q not found", arg)
				return nil
			}

			return err
		}

	case "c":
		err := r.todoStore.Create(arg)
		if err != nil {
			if todo.IsNotFoundErr(err) {
				fmt.Printf("TodoItem %q not found", arg)
				return nil
			}

			return err
		}

	case "h":
		r.l.Println("invalid input: help command expects zero arguments")
		fmt.Println("help command expects zero arguments")
	default:
		r.l.Printf("unknown command: %q\n", instruction)
		fmt.Printf("unknown command: %q\n", instruction)
	}

	return nil
}

func (r Routes) writeTodoList(list []todo.Item) {
	for _, t := range list {
		fmt.Fprintf(r.tw, "%s\t%s\t%t\n", t.UID, t.Task, t.Checked)
	}

	r.tw.Flush()
	fmt.Println()
}
