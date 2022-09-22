package todo

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/alvinmatias69/golang-onboarding/internal/model/todo"
)

type operation int

// define possible operation input
const (
	list operation = iota
	add
	mark
)

// mapping input string to enum
var opsMap = map[string]operation{
	"LIST": list,
	"ADD":  add,
	"MARK": mark,
}

// Parse params and handle it accordingly
func (h *Handler) Parse(params string) error {
	if len(params) < 1 {
		return errors.New("argument required")
	}

	var (
		args = strings.Split(params, " ")
		rest = strings.Join(args[1:], " ")
		err  error
	)

	ops, ok := opsMap[strings.Trim(args[0], "\n")]
	if !ok {
		printHelp()
		return nil
	}

	switch ops {
	case list:
		err = h.list(rest)
	case add:
		err = h.add(rest)
	case mark:
		err = h.mark(rest)
	}

	return err
}

func (h *Handler) list(params string) error {
	fmt.Println("Listing all todo list item")
	filter := todo.TodoFilter{}
	for _, arg := range strings.Split(params, " ") {
		if id, err := strconv.ParseUint(arg, 10, 64); err == nil {
			filter.ID = id
			continue
		}
		if isDone, err := strconv.ParseBool(arg); err == nil {
			filter.IsDone = &isDone
		}
	}
	res, err := h.usecase.List(context.Background(), filter)
	if err != nil {
		return err
	}

	for _, item := range res {
		printTodo(item)
	}

	return nil
}

// TODO: implement this
func (h *Handler) add(params string) error {
	fmt.Println("Adding a todo list item")
	return nil
}

// TODO: implement this
func (h *Handler) mark(params string) error {
	fmt.Println("Mark a todo list item as done")
	return nil
}
