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
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add command invoked")

		if len(args) > 0 {
			desc := args[0]
			fmt.Println(desc)
			lastID := getLastID()
			if lastID == -1 {
				panic("Error getting last id")
			}

			todo := &Todo{
				id:        lastID + 1,
				desc:      desc,
				createdAt: time.Now(),
			}

			file, err := os.OpenFile("todos.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
			if err != nil {
				fmt.Println("Error creating file: ", err)
				return
			}
			defer file.Close()

			writer := csv.NewWriter(file)
			defer writer.Flush()

			newTodoSlice := []string{
				strconv.Itoa(todo.id),
				todo.desc,
				todo.createdAt.String(),
				strconv.FormatBool(todo.isCompleted),
			}

			err = writer.Write(newTodoSlice)
			if err != nil {
				fmt.Println("Error wrtiing file: ", err)
				return
			}

			fmt.Println("Todo addec successfully")

		} else {
			fmt.Println("No description provided")
		}
	},
}

func getLastID() int {
	file, err := os.OpenFile("todos.csv", os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return -1
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var lastLine []string

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		lastLine = row
	}

	if lastLine == nil {
		// Create Headers
		writer := csv.NewWriter(file)
		defer writer.Flush()

		if err := writer.Write([]string{
			"ID",
			"Description",
			"CreatedAt",
			"IsCompleted",
		}); err != nil {
			fmt.Println("Error while creating headers: ", err)
		}

		return 0
	}

	lastID, err := strconv.Atoi(lastLine[0])
	if err != nil {
		fmt.Println("Error parsing id: ", err)
		return -1
	}

	return lastID
}
