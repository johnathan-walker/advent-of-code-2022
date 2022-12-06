package cmd

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func createFile(parts ...string) io.ReadCloser {
	inputFilePath := filepath.Join(parts...)
	f, fErr := os.OpenFile(inputFilePath, os.O_RDONLY, os.ModePerm)
	if fErr != nil {
		log.Printf("error opening file: %v\n", fErr)
		os.Exit(1)
	}
	return f
}
