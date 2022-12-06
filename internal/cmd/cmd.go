package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Verbose bool
var SourceDir string

var rootCmd = &cobra.Command{
	Use:   "advent-of-code-2022",
	Short: "advent-of-code-2022 runs the advent of code solutions.",
}

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Verbose output")
	rootCmd.PersistentFlags().StringVarP(&SourceDir, "dir", "d", "", "Source directory to read from")
	rootCmd.MarkPersistentFlagRequired("dir")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
