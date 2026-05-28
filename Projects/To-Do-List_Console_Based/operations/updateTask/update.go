package updatetask

import (
	"encoding/json"
	"fmt"
	"os"
	userinput "to-do-list/User-Input"
	store "to-do-list/task-store"
)

func UpdateTask() {
	for {
		fmt.Print("Enter Task ID to mark complete (0 to exit): ")
		id := -1
		id = userinput.UserInputInt()

		if id == 0 {
			fmt.Println("Leaving update menu...")
			return
		}

		tasks := store.LoadTasks()
		found := false

		for i, t := range tasks {
			if t.ID == id {
				found = true
				if t.Completed {
					fmt.Println("Task already marked as complete. Try another ID.")
					break
				}
				tasks[i].Completed = true
				fmt.Println("Task marked as complete!")

				// Save updated tasks back to file
				data, err := json.MarshalIndent(tasks, "", "  ")
				if err != nil {
					fmt.Println("Error marshaling tasks:", err)
					return
				}
				err = os.WriteFile(store.FilePath(), data, 0644)
				if err != nil {
					fmt.Println("Error writing file:", err)
				}
				return
			}
		}

		if !found {
			fmt.Println("Task ID not found. Try again.")
		}
	}
}
