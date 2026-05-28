package addtask

import (
	"fmt"
	userinput "to-do-list/User-Input"
	types "to-do-list/model"
	store "to-do-list/task-store"
)

func AddTask() {
	fmt.Print("Enter Task Name: ")
	var name string
	name = userinput.UserInputString()

	task := types.Add_Task{
		ID:        store.GetNextID(),
		Name:      name,
		Completed: false,
	}

	// Save task to file
	store.SaveTask(task)

	fmt.Println("Task added successfully!")
}
