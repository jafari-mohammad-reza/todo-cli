package todo

import (
	"errors"
	"fmt"
	"github.com/jafari-mohammad-reza/todo-cli/utils"
	"time"
)

type Item struct {
	Title     string
	IsDone    bool
	CreatedAt time.Time
	DoneAt    time.Time
}
type Todos []Item

func (t *Todos) AddTask(task string) {
	newTodo := Item{Title: task, IsDone: false, CreatedAt: time.Now(), DoneAt: time.Time{}}
	*t = append(*t, newTodo)
	_, err := utils.AppendDataToFile(utils.GetSaveFilePath(), newTodo)
	if err != nil {
		return
	}
}
func (t *Todos) CompleteTask(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	task := ls[index-1]
	compareKey := "Title"
	key := task.Title

	mapSlice := todosToMapSlice(ls)
	utils.UpdateJSONProperty(utils.GetSaveFilePath(), mapSlice, key, true, compareKey, func(item *map[string]interface{}, val interface{}) {
		(*item)["IsDone"] = val.(bool)
		(*item)["DoneAt"] = time.Now().Format(time.RFC3339)
	})

	return nil
}

func todosToMapSlice(todos Todos) []map[string]interface{} {
	mapSlice := make([]map[string]interface{}, len(todos))
	for i, todo := range todos {
		mapSlice[i] = map[string]interface{}{
			"Title":  todo.Title,
			"IsDone": todo.IsDone,
			"DoneAt": todo.DoneAt,
		}
	}
	return mapSlice
}

func (t *Todos) DeleteTask(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}
	utils.Remove(ls, index)
	return nil
}
func (t *Todos) Print() {
	for _, item := range *t {
		fmt.Printf("Title : %v IsDone : %v  \n", item.Title, item.IsDone)
	}
}
