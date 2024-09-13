package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteTodo)
}

var deleteTodo = &cobra.Command{
	Use: "del",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete todo invoked")
	},
}
