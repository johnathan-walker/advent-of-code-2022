package cmd

import (
	"fmt"
	"log"
	"os"

	rockpaperscissors "github.com/johnathan-walker/advent-of-code-2022/internal/rockPaperScissors"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dayTwo)
	dayTwo.AddCommand(d2s1)
	dayTwo.AddCommand(d2s2)
}

var dayTwo = &cobra.Command{
	Use:   "dayTwo",
	Short: "Runs dayTwo solution",
}

var d2s1 = &cobra.Command{
	Use:   "starOne",
	Short: "Runs daytwo starOne solution",
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

var d2s2 = &cobra.Command{
	Use:   "starTwo",
	Short: "Runs dayOne starTwo solution",
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
