package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vtigo/cli-task-manager/internal/db"
	"github.com/vtigo/cli-task-manager/internal/handlers"
	"github.com/vtigo/cli-task-manager/internal/models"
)

func main() {
	taskManager := models.NewTaskManager()
	storage := db.NewFileStorage("storage")
	handler := handlers.NewTaskHandler(taskManager, storage)

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Failed to load tasks: ", err)
		os.Exit(1)
	}

	for _, t := range tasks {
		taskManager.Tasks = append(taskManager.Tasks, t)
	}

	rootCmd := &cobra.Command{
		Use: "ctm",
		Short: "Task list application",
		Long: "Task list CLI application made for GOLANG learning purpose",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to CTM! Use --help to see available commands")
		},
	}

	listCmd := &cobra.Command{
		Use: "list",
		Short: "List all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			handler.HandleListTasks()
		},
	}

	createCmd := &cobra.Command{
		Use: "task",
		Short: "Create a new task",
		Run: func(cmd *cobra.Command, args []string) {
			handler.HandleCreateTask(args[0])
		},
	}

	completeCmd := &cobra.Command{
		Use: "tick",
		Short: "Complete a task",
		Run: func(cmd *cobra.Command, args []string) {
			intInput, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("failed to convert task index argument to integer")
			}
			
			taskIndex := intInput - 1
			handler.HandleMarkAsCompleted(taskIndex)
		},
	}

	deleteCmd := &cobra.Command{
		Use: "delete",
		Short: "Delete a task",
		Run: func(cmd *cobra.Command, args []string) {
			intInput, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("failed to convert task index argument to integer")
			}
	
			taskIndex := intInput - 1
			handler.HandleDeleteTask(taskIndex)
		},
	}

	rootCmd.AddCommand(
		listCmd,
		createCmd,
		completeCmd,
		deleteCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

