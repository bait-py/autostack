package cmd

import (
	"github.com/bait-py/autostack/internal/stack"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available stacks",
	Run: func(cmd *cobra.Command, args []string) {
		stack.ListStacks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
