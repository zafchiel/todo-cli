package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Todo struct {
	id          int
	desc        string
	createdAt   time.Time
	isCompleted bool
}

var rootCmd = &cobra.Command{
	Use:   "tc",
	Short: "tc (todo-cli) manages todos with files",
	Long: `A simple CLI application for managing your todos using files.
						It allows you to create, list, update, and delete todos
						easily from your terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		// If no subcommand is provided, print help
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
