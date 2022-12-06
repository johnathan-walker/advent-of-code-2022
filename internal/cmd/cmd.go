package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "advent-of-code-2022",
	Short: "advent-of-code-2022 runs the advent of code solutions.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
