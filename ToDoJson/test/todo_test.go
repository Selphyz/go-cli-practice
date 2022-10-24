package test

import (
	"ToDoJson/cmd/todo"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, l[0].Task)
	}
}
func TestComplete(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q, go %q instead", taskName, l[0].Task)
	}
	if l[0].Done {
		t.Errorf("New task should not be completed")
	}
	err := l.Complete(1)
	if err != nil {
		t.Errorf("Error completing the task")
	}
	if !l[0].Done {
		t.Errorf("New task should be completed")
	}
}
func TestDelete(t *testing.T) {
	l := todo.List{}
	tasks := []string{
		"New task 1",
		"New task 2",
		"New task 3",
	}
	for i, v := range tasks {
		l.Add(v)
		if l[i].Task != tasks[i] {
			t.Errorf("Expected %q, got %q instead.", tasks[i], l[i].Task)
		}
	}
	err := l.Delete(2)
	if err != nil {
		t.Errorf("Error deleting task number 2")
	}
	if len(l) != 2 {
		t.Errorf("Expected %q, got %q instead", tasks[2], l[1].Task)
	}
}
func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}
	taskName := "New Task"
	l1.Add(taskName)
	if l1[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, l1[0].Task)
	}
	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatalf("Error deleting temp file: %s", err)
		}
	}(tf.Name())
	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}
	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}
	l2.Add(taskName)
	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q", l1[0].Task, l2[0].Task)
	}
}
