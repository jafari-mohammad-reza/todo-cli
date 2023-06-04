package main

import (
	"flag"
	"fmt"
	todos "github.com/jafari-mohammad-reza/todo-cli/todo"
	"github.com/jafari-mohammad-reza/todo-cli/utils"
	"log"
	"os"
)

func main() {
	utils.CreateDataFile()
	add := flag.Bool("add", false, "add new todo")
	empty := flag.Bool("empty", false, "remove all data")
	complete := flag.Int("complete", 0, "complete a task")
	list := flag.Bool("list", false, "get list of tasks")
	flag.Parse()
	todo := &todos.Todos{}
	data, err := utils.ReadFromDataFile[[]todos.Item](utils.GetSaveFilePath())
	if err != nil {
		fmt.Println(err)
		log.Fatal("Failed to load data")
	}
	*todo = data
	switch {
	case *add:
		task, err := utils.GetInput(os.Stdin, flag.Args()...)
		if err != nil {
			log.Fatal("failed to read input")
		}
		todo.AddTask(task)
	case *empty:
		utils.RemoveAllDataFromFile(utils.GetSaveFilePath())
	case *complete > 0:
		completeErr := todo.CompleteTask(*complete)
		if completeErr != nil {
			return
		}
	case *list:
		todo.Print()
	default:
		log.Fatal("Invalid Command")
	}
}
