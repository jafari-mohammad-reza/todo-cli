package todo

import (
	"errors"
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
	task.IsDone = true
	task.DoneAt = time.Now()
	_, err := utils.AppendDataToFile(utils.GetSaveFilePath(), task)
	if err != nil {
		return err
	}
	return nil
}
func (t *Todos) DeleteTask(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}
	utils.Remove(ls, index)
	return nil
}
