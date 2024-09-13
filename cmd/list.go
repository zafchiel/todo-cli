package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listTodos)
}

var listTodos = &cobra.Command{
	Use: "all",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete todo invoked")
	},
}
