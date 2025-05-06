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

type FileStorage struct {
	StoragePath string
}

func NewFileStorage(storagePath string) *FileStorage {
	return &FileStorage{
		StoragePath: storagePath,
	}
}

func (s *FileStorage) SaveTaskList(tl *models.TaskList) error {
	if err := os.MkdirAll(s.StoragePath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create storage directory: %w", err)
	}
	
	jsonData, err := json.Marshal(tl)
	if err != nil {
		return fmt.Errorf("failed to marshal task list: %w", err)
	}
	
	path := filepath.Join(s.StoragePath, fmt.Sprintf("%s.json", tl.Name))
	if err := os.WriteFile(path, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	
	return nil
}

func (s *FileStorage) LoadAllTaskLists() ([]*models.TaskList, error) {
	taskLists := []*models.TaskList{}
	
	if _, err := os.Stat(s.StoragePath); os.IsNotExist(err) {
		return taskLists, nil
	}
	
	root := os.DirFS(s.StoragePath)
	taskListFiles, err := fs.Glob(root, "*.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read storage directory: %w", err)
	}
	
	for _, filename := range taskListFiles {
		path := filepath.Join(s.StoragePath, filename)
		file, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
		}
		
		name := strings.Split(filename, ".")[0]
		
		tl := &models.TaskList{
			Name: name,
		}
		
		if err := json.Unmarshal(file, tl); err != nil {
			return nil, fmt.Errorf("failed to unmarshal task list from %s: %w", filename, err)
		}
		
		taskLists = append(taskLists, tl)
	}

	// TODO: Find main task list
	// if it doesnt exist, create it
	
	return taskLists, nil
}
