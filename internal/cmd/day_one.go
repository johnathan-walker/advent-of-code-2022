package cmd

import (
	"fmt"
	"log"
	"os"

	dayone "github.com/johnathan-walker/advent-of-code-2022/internal/dayOne"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dayOne)
	dayOne.AddCommand(starOne)
	dayOne.AddCommand(starTwo)
}

var dayOne = &cobra.Command{
	Use:   "dayOne",
	Short: "Runs dayOne solution",
}

var starOne = &cobra.Command{
	Use:   "starOne",
	Short: "Runs dayOne starOne solution",
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

		cals, err := dayone.RunStarOne(f)
		if err != nil {
			log.Printf("could not run code: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(cals)
	},
}

var starTwo = &cobra.Command{
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

		topThree, err := dayone.RunStarTwo(f)
		if err != nil {
			log.Printf("could not run code: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(topThree[0] + topThree[1] + topThree[2])
	},
}
