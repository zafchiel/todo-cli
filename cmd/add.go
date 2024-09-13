package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

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
			todo := &Todo{desc: desc}

			file, err := os.OpenFile("todos.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
			if err != nil {
				fmt.Println("Error creating file: ", err)
				return
			}
			defer file.Close()

			writer := csv.NewWriter(file)
			defer writer.Flush()

			err = writer.Write([]string{strconv.Itoa(todo.id), todo.desc, todo.createdAt.String(), strconv.FormatBool(todo.isCompleted)})
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
