package cmd

import (
	"fmt"
	"log"
	"os"

	rockpaperscissors "github.com/johnathan-walker/advent-of-code-2022/internal/rockPaperScissors"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(d2Cmd)
	d2Cmd.AddCommand(d2s1Cmd)
	d2Cmd.AddCommand(d2s2Cmd)
}

var d2Cmd = &cobra.Command{
	Use:   "d2",
	Short: "Runs the d2 solution.",
}

var d2s1Cmd = &cobra.Command{
	Use:   "s1",
	Short: "Runs the d2 s1 solution.",
	Run: func(cmd *cobra.Command, args []string) {
		pwd, err := os.Getwd()
		if err != nil {
			log.Printf("error determining pwd: %s\n", err.Error())
			os.Exit(1)
		}
		f := createFile(pwd, SourceDir, "input.txt")
		defer func() {
			if err := f.Close(); err != nil {
				log.Printf("error closing file: %s\n", err.Error())
				os.Exit(1)
			}
		}()

		totalScore, err := rockpaperscissors.PartOne(f)
		if err != nil {
			log.Printf("could not run code: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(totalScore)
	},
}

var d2s2Cmd = &cobra.Command{
	Use:   "s2",
	Short: "Runs the d2 s2 solution.",
	Run: func(cmd *cobra.Command, args []string) {
		pwd, err := os.Getwd()
		if err != nil {
			log.Printf("error determining pwd: %s\n", err.Error())
			os.Exit(1)
		}
		f := createFile(pwd, SourceDir, "input.txt")
		defer func() {
			if err := f.Close(); err != nil {
				log.Printf("error closing file: %s\n", err.Error())
				os.Exit(1)
			}
		}()

		totalScore, err := rockpaperscissors.PartTwo(f)
		if err != nil {
			log.Printf("could not run code: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(totalScore)
	},
}
