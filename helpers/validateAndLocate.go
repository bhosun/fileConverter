package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ValidateAndLocateFile(fileName string) (string, []string) {
	splittedFileName := strings.Split(fileName, ".")
	if splittedFileName[1] != "txt" {
		fmt.Println("Kindly input the right fileName that ends with .txt!")
	}

	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error", err)
		return "", nil
	}

	fileFound := false
	filePath := ""

	fileSearchErr := filepath.Walk(workingDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == fileName {
			filePath = path
			fileFound = true
		}

		return nil
	})

	if fileSearchErr != nil {
		fmt.Println("Error:", fileSearchErr)
		return "", nil
	}

	if !fileFound {
		fmt.Println("File not found.")
		return "", nil
	}

	return filePath, splittedFileName
}
