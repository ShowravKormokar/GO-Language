package deletetask

import (
	"encoding/json"
	"fmt"
	"os"
	userinput "to-do-list/User-Input"
	"to-do-list/model"
	store "to-do-list/task-store"
)

// DeleteTask removes a task by ID
func DeleteTask() {
	fmt.Print("Enter Task ID to delete (0 to cancel): ")
	id := -1
	id = userinput.UserInputInt()

	if id == 0 {
		fmt.Println("Delete cancelled.")
		return
	}

	tasks := store.LoadTasks()
	newTasks := []model.Add_Task{}
	found := false

	for _, t := range tasks {
		if t.ID == id {
			found = true
			fmt.Printf("Task ID %d deleted successfully!\n", id)
			continue // skip this task
		}
		newTasks = append(newTasks, t)
	}

	if !found {
		fmt.Println("Task ID not found.")
		return
	}

	// Save updated list back to file
	data, err := json.MarshalIndent(newTasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling tasks:", err)
		return
	}
	err = os.WriteFile(store.FilePath(), data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}
