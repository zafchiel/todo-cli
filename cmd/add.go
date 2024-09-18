package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addTodo)
}

var addTodo = &cobra.Command{
	Use:   "add <description>",
	Short: "Add a new todo item",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		desc := args[0]
		lastID := getLastID()
		if lastID == -1 {
			fmt.Fprintln(os.Stderr, "Error getting last ID")
			return
		}

		todo := &Todo{
			id:        lastID + 1,
			desc:      desc,
			createdAt: time.Now(),
		}

		if err := writeTodo(todo); err != nil {
			fmt.Fprintf(os.Stderr, "Error adding todo: %v\n", err)
			return
		}

		fmt.Println("Todo added successfully, ID - ", todo.id)
	},
}

func writeTodo(todo *Todo) error {
	file, err := os.OpenFile("todos.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.Write([]string{
		strconv.Itoa(todo.id),
		todo.desc,
		todo.createdAt.Format(time.RFC3339),
		strconv.FormatBool(todo.isCompleted),
	})
}

func getLastID() int {
	file, err := os.OpenFile("todos.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		return -1
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		return -1
	}

	if len(rows) == 0 {
		// Create Headers
		writer := csv.NewWriter(file)
		defer writer.Flush()

		if err := writer.Write([]string{"ID", "Description", "CreatedAt", "IsCompleted"}); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating headers: %v\n", err)
			return -1
		}
		return 0
	}

	lastID, err := strconv.Atoi(rows[len(rows)-1][0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing ID: %v\n", err)
		return -1
	}

	return lastID
}
