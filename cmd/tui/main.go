package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/vtigo/cli-task-manager/internal/db"
	"github.com/vtigo/cli-task-manager/internal/handlers"
	"github.com/vtigo/cli-task-manager/internal/models"
)

type model struct {
	tasks  	  []*models.Task
	cursor    int
	handler   *handlers.TaskHandler
}

func initialModel() model {
	// Initiate managers (tasks, storage)
	taskManager := models.NewTaskManager()
	storage := db.NewFileStorage("storage")
	handler := handlers.NewTaskHandler(taskManager,storage)
	
	// Load tasks from storage and add them to taskManager.Tasks
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Failed to load tasks: ", err)
		os.Exit(1)
	}
	for _, t := range tasks {
		taskManager.Tasks = append(taskManager.Tasks, t)
	}
	
	return model{
		tasks: taskManager.Tasks,
		handler: handler,
	}
}

func (m model) Init() tea.Cmd {
	// Initiate model with no commands
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {

        // These keys should exit the program.
        case "ctrl+c", "q":
            return m, tea.Quit

        // The "up" and "k" keys move the cursor up
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        // The "down" and "j" keys move the cursor down
        case "down", "j":
            if m.cursor < len(m.tasks)-1 {
                m.cursor++
            }

		// Toggle the completeness of the task underneath the cursor
        case "enter", " ":
			if err := m.handler.HandleToggleTaskCompleteness(m.cursor); err != nil {
				fmt.Println("failed to mark task as completed:", err)
				os.Exit(1)
			}

		// Delete the task underneath the cursor
		case "d":
			if err := m.handler.HandleDeleteTask(m.cursor); err != nil {
				fmt.Println("failed to deled task:", err)
				os.Exit(1)
			}
        }
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m model) View() string {
    // The header
    s := "How are you today?\n\n"

    // Iterate over our tasks
    for i, task := range m.tasks {

        // Is the cursor pointing at this task?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
        }

        // Is this task completed?
        checked := " " // not selected
        if m.tasks[i].Completed {
            checked = "x" // selected!
        }

        // Render the row
        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, task.Name)
    }

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    return s
}

func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
