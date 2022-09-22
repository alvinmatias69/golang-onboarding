package todo

import (
	"fmt"

	"github.com/alvinmatias69/golang-onboarding/internal/model/todo"
)

func printTodo(item todo.Todo) {
	fmt.Println("===")
	fmt.Printf("ID\t\t: %v\n", item.ID)
	fmt.Printf("Task\t\t: %v\n", item.Task)
	fmt.Printf("IsDone\t\t: %v\n", item.IsDone)
	fmt.Printf("CreatedAt\t: %v\n", item.CreatedAt)
	fmt.Printf("UpdatedAt\t: %v\n", item.UpdatedAt)
	fmt.Println("===")
}

func printHelp() {
	fmt.Println(`Usage help:

HELP
	Print this page
LIST <params>
	Display list of todo items, accepting optional params which can be boolean (is done) or uint (id).
	[example: LIST true 15]
ADD <task_name>
	Add new todo item.
	[example: ADD masak air]
MARK <task_id>
	Mark task as done.
	[example: MARK 15]`)
}
