package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	types "to-do-list/model"
)

var filePath = filepath.Join(".", "task-store", "task.json")

func FilePath() string {
	return filePath
}

// SaveTask appends a new task to task.json
func SaveTask(task any) {
	tasks := LoadTasks()
	tasks = append(tasks, task.(types.Add_Task))

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling tasks:", err)
		return
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

// GetNextID returns the next unique ID based on existing tasks
func GetNextID() int {
	tasks := LoadTasks()
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

// LoadTasks reads all tasks from task.json
func LoadTasks() []types.Add_Task {
	var tasks []types.Add_Task

	// Ensure file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		_ = os.WriteFile(filePath, []byte("[]"), 0644)
		return tasks
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return tasks
	}

	if len(data) == 0 {
		return tasks
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error unmarshaling tasks:", err)
		return tasks
	}

	return tasks
}
