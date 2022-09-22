package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	todoHandler "github.com/alvinmatias69/golang-onboarding/internal/handler/todo"
	todoRepository "github.com/alvinmatias69/golang-onboarding/internal/repository/todo"
	todoUsecase "github.com/alvinmatias69/golang-onboarding/internal/usecase/todo"
)

func main() {
	fmt.Println("===========================")
	fmt.Println(" Welcome to simple project")
	fmt.Println("===========================")

	var (
		todoRepo = todoRepository.New()
		todoUc   = todoUsecase.New(todoRepo)
		todoHnd  = todoHandler.New(todoUc)
		reader   = bufio.NewReader(os.Stdin)
	)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if err = todoHnd.Parse(line); err != nil {
			log.Printf("error: %s\n", err.Error())
		}
	}
}
