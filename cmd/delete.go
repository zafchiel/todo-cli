package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteTodo)
}

var deleteTodo = &cobra.Command{
	Use:   "del <todo-id>",
	Short: "Delete a todo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete todo invoked")
		id := args[0]

		file, err := os.OpenFile("todos.csv", os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error opening a file: ", err)
			return
		}

		defer file.Close()

		reader := csv.NewReader(file)
		var newRows [][]string

		for {
			row, err := reader.Read()
			if err != nil {
				if err == io.EOF {
					// End of file
					break
				}
				fmt.Fprint(os.Stderr, "Error reading csv, ", err)
				return
			}
			if row[0] == id {
				// Dont add to new rows
			} else {
				newRows = append(newRows, row)
			}
		}

		file.Seek(0, 0)
		file.Truncate(0)

		writer := csv.NewWriter(file)
		if err := writer.WriteAll(newRows); err != nil {
			fmt.Fprint(os.Stderr, "Error writing to CSV file: ", err)
			return
		}

		fmt.Printf("Task %s deleted\n", id)
	},
}
