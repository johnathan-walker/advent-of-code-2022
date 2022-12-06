package cmd

import (
	"fmt"
	"log"
	"os"

	dayone "github.com/johnathan-walker/advent-of-code-2022/internal/countingcalories"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(d1Cmd)
	d1Cmd.AddCommand(d1s1Cmd)
	d1Cmd.AddCommand(d1s2Cmd)
}

var d1Cmd = &cobra.Command{
	Use:   "d1",
	Short: "Runs the d1 solution.",
}

var d1s1Cmd = &cobra.Command{
	Use:   "s1",
	Short: "Runs the d1 s1 solution.",
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

		cals, err := dayone.PartOne(f)
		if err != nil {
			log.Printf("could not run code: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(cals)
	},
}

var d1s2Cmd = &cobra.Command{
	Use:   "s2",
	Short: "Runs the d1 s2 solution.",
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

		topThree, err := dayone.PartTwo(f)
		if err != nil {
			log.Printf("could not run code: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(topThree[0] + topThree[1] + topThree[2])
	},
}
