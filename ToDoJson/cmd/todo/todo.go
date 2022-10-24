package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}
type List []Item

func (l *List) Add(task string) {
	t := Item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()
	return nil
}
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}
func (l List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, &l)
}
func (l List) Save(fileName string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, js, 0644)
}
