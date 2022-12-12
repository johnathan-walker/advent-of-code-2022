package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/johnathan-walker/advent-of-code-2022/internal/rucksack"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(d3Cmd)
	d3Cmd.AddCommand(d3s1Cmd)
}

var d3Cmd = &cobra.Command{
	Use:   "d3",
	Short: "Runs the d3 solution.",
}

var d3s1Cmd = &cobra.Command{
	Use:   "s1",
	Short: "Runs the d3 s1 solution.",
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

		totalScore, err := rucksack.PartOne(f)
		if err != nil {
			log.Printf("could not run code: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(totalScore)

	},
}
