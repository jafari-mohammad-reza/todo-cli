package todo

import (
	"errors"
	"github.com/jafari-mohammad-reza/todo-cli/utils"
	"time"
)

type item struct {
	Title     string
	IsDone    bool
	CreatedAt time.Time
	DoneAt    time.Time
}
type Todos []item

func (t *Todos) AddTask(task string) {
	newTodo := item{Title: task, IsDone: false, CreatedAt: time.Now(), DoneAt: time.Time{}}
	*t = append(*t, newTodo)
}
func (t *Todos) CompleteTask(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}
	task := ls[index-1]
	task.IsDone = true
	task.DoneAt = time.Now()
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
