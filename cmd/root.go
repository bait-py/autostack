package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "autostack",
	Short: "AutoStack CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("AutoStack CLI. Use -h for help.")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
