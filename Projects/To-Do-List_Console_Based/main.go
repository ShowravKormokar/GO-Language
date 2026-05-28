package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
	menu "to-do-list/Menu"
	userinput "to-do-list/User-Input"
	addOper "to-do-list/operations/addTask"
	deleteOper "to-do-list/operations/delete-task"
	updateOper "to-do-list/operations/updateTask"
	viewOper "to-do-list/operations/view-tasks"
)

// ClearScreen clears the terminal screen based on the OS
func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Windows command
	} else {
		cmd = exec.Command("clear") // Unix/Linux/Mac command
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	for {
		menu.Menu()
		fmt.Print("Enter your choice: ")
		input := userinput.UserInputInt()

		switch input {
		case 1:
			fmt.Println("========= Add task =========")
			addOper.AddTask()
		case 2:
			fmt.Println("========= View Tasks =========")
			viewOper.ViewTasks()
		case 3:
			fmt.Println("========= Mark Task As Done =========")
			updateOper.UpdateTask()
		case 4:
			fmt.Println("========= Delete Task =========")
			deleteOper.DeleteTask()
		case 0:
			fmt.Println("Exiting...")
			time.Sleep(1 * time.Second)
			ClearScreen()
			fmt.Println("Exited successfully.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
