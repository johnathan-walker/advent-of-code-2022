package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	dayone "github.com/johnathan-walker/advent-of-code-2022/internal/dayOne"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dayOne)
}

var dayOne = &cobra.Command{
	Use:   "dayOne",
	Short: "Runs dayOne solution",
	Run: func(cmd *cobra.Command, args []string) {
		pwd, _ := os.Getwd()

		inputFilePath := filepath.Join(pwd, SourceDir, "input.txt")
		f, fErr := os.OpenFile(inputFilePath, os.O_RDONLY, os.ModePerm)
		defer func() {
			if err := f.Close(); err != nil {
				log.Printf("error closing file %s: %s\n", inputFilePath, err.Error())
				os.Exit(1)
			}
		}()
		if fErr != nil {
			log.Printf("error opening file: %v\n", fErr)
			os.Exit(1)
		}

		id, err := dayone.Run(f)
		if err != nil {
			log.Printf("could not run code: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(id)
	},
}
