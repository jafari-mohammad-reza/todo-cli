package main

import (
	"flag"
	"fmt"
	todos "github.com/jafari-mohammad-reza/todo-cli/todo"
	"github.com/jafari-mohammad-reza/todo-cli/utils"
	"log"
)

func main() {
	utils.CreateDataFile()
	add := flag.Bool("add", false, "add new todo")
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
		todo.AddTask("test")
	}
}
