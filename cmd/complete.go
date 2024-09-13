package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completeTodo)
}

var completeTodo = &cobra.Command{
	Use: "done",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete todo invoked")
	},
}
