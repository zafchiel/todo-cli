package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completeTodo)
}

var completeTodo = &cobra.Command{
	Use:   "done <task-id>",
	Short: "Mark task as completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		file, err := os.OpenFile("todos.csv", os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error opening a file: ", err)
		}
		defer file.Close()

		reader := csv.NewReader(file)

		var newRows [][]string

		for {
			row, err := reader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Fprint(os.Stderr, "Error reading csv, ", err)
			}

			if row[0] == id {
				fmt.Fprint(os.Stdout, "FOUND TODO, ", id)
				row[3] = "true"
			}

			newRows = append(newRows, row)
		}

		file.Seek(0, 0)
		file.Truncate(0)

		writer := csv.NewWriter(file)
		if err := writer.WriteAll(newRows); err != nil {
			fmt.Fprint(os.Stderr, "Error writing to CSV file: ", err)
			return
		}

		fmt.Printf("Task %s marked as completed\n", id)
	},
}
