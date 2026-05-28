package viewtasks

import (
	"fmt"
	store "to-do-list/task-store"
)

// ViewTasks loads all tasks and prints them
func ViewTasks() {
	tasks := store.LoadTasks()

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("Your Tasks:")
	fmt.Println("-----------")
	for _, t := range tasks {
		var status string
		if t.Completed == true {
			status = "✅ Completed"
		} else if t.Completed == false {
			status = "❌ Not Completed"
		}
		fmt.Printf("ID: %d | Name: %s | Status: %s\n", t.ID, t.Name, status)
	}
}
