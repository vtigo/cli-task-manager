package db

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/vtigo/cli-task-manager/internal/models"
)

func HandleSaveLists() {
	for _, tl := range models.TaskLists{
		tl.SaveFile()
	}
}

func HandleLoadTaskLists() []*models.TaskList {
	models.TaskLists = []*models.TaskList{}

	root := os.DirFS("./storage")
	taskListFiles, err := fs.Glob(root, "*.json")
	if err != nil {
		fmt.Println("Failed to read storage directory")
		return models.TaskLists
	}

	for i := range taskListFiles {
		path := filepath.Join("storage", taskListFiles[i])
		file, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("failed to read file:", err)
			continue
		}
	
		name := strings.Split(taskListFiles[i], ".")[0]
		tl := models.NewTaskList(name)

		err = json.Unmarshal(file, &tl)
		if err != nil {
			fmt.Println("failed to unmarshal task list:", err)
			continue
		}

		for j := range tl.Tasks {
			models.Tasks = append(models.Tasks, &tl.Tasks[j])
		}
	}
	
	models.TaskListId = len(models.TaskLists)
	models.TaskId = len(models.Tasks)
	
	return models.TaskLists
}

