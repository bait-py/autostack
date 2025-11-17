package cmd

import (
	"fmt"

	"github.com/bait-py/autostack/internal/stack"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [stack]",
	Short: "Create a stack",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		stackName := args[0]
		fmt.Printf("Creating stack: %s\n", stackName)
		return stack.Create(stackName)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
