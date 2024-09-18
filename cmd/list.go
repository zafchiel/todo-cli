package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listTodos)
	listTodos.Flags().BoolVarP(&All, "all", "a", false, "List all todos")
}

var All bool

var listTodos = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.OpenFile("todos.csv", os.O_RDONLY, 0644)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error opening a file: ", err)
			return
		}
		defer file.Close()

		// Check if the file is empty
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error getting file info:", err)
			return
		}

		if fileInfo.Size() == 0 {
			fmt.Println("No todos found.")
			return
		}

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 8, 0, '\t', 0)
		defer w.Flush()

		reader := csv.NewReader(file)

		// Display heder
		header, err := reader.Read()
		if err != nil {
			fmt.Fprint(os.Stderr, "Error reading header from csv, ", err)
			return
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", header[0], header[1], header[2], header[3])

		// Print rows
		for {
			row, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Fprint(os.Stderr, "Error reading csv, ", err)
				return
			}

			if All || row[3] == "false" {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", row[0], row[1], row[2], row[3])
			}

		}

	},
}
