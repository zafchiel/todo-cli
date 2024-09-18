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
	listTodos.Flags().BoolVarP(&listAll, "all", "a", false, "List all todos including completed ones")
}

var listAll bool

var listTodos = &cobra.Command{
	Use:   "list",
	Short: "List uncompleted todos",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open("todos.csv")
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

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 3, '\t', 0)
		defer w.Flush()

		reader := csv.NewReader(file)

		// Display heder
		header, err := reader.Read()
		if err != nil {
			fmt.Fprint(os.Stderr, "Error reading header from csv, ", err)
			return
		}
		printTab(w, header)

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

			if listAll || row[3] == "false" {
				printTab(w, row)
			}

		}

	},
}

func printTab(w *tabwriter.Writer, row []string) {
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", row[0], row[1], row[2], row[3])
}
