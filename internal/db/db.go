package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

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

func (s *FileStorage) SaveTasks(tasks []*models.Task) error {
	if err := os.MkdirAll(s.StoragePath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create storage directory: %w", err)
	}
	
	jsonData, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("failed to marshal task list: %w", err)
	}
	
	path := filepath.Join(s.StoragePath, fmt.Sprintf("%s.json", "tasks"))
	if err := os.WriteFile(path, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	
	return nil
}

func (s *FileStorage) LoadTasks() ([]*models.Task, error) {
	tasks := []*models.Task{}

	path := filepath.Join(s.StoragePath, "tasks.json")
	
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return tasks, nil
	}
	
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file tasks.json: %w", err)
	}
	
	if err = json.Unmarshal(file, &tasks); err != nil {
		return nil, fmt.Errorf("failed to unmarshal task list from tasks.json: %w", err)
	}

	return tasks, nil
}
