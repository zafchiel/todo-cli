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
		id := args[0]

		file, err := os.OpenFile("todos.csv", os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error opening a file: ", err)
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)
		var newRows [][]string
		var found bool

		for {
			row, err := reader.Read()
			if err == io.EOF {
				// End of file
				break
			}
			if err != nil {
				fmt.Fprint(os.Stderr, "Error reading csv, ", err)
				return
			}
			if row[0] == id {
				// Dont add to new rows
				found = false
			} else {
				newRows = append(newRows, row)
			}
		}

		if !found {
			fmt.Fprint(os.Stderr, "Todo not found\n")
			return
		}

		if _, err := file.Seek(0, 0); err != nil {
			fmt.Fprint(os.Stderr, "Error seeking to start of file: ", err)
			return
		}
		if err := file.Truncate(0); err != nil {
			fmt.Fprint(os.Stderr, "Error truncating file: ", err)
			return
		}

		writer := csv.NewWriter(file)
		if err := writer.WriteAll(newRows); err != nil {
			fmt.Fprint(os.Stderr, "Error writing to CSV file: ", err)
			return
		}

		fmt.Printf("Task %s deleted\n", id)
	},
}
