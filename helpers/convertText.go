package helpers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func ConvertText2Json(filePath string, splittedFileName []string) {
	// Open file to be converted
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Write values into new Map and convert to json
	scanner := bufio.NewScanner(file)
	newMap := make(map[int]string)
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		key := lineCount
		newMap[key] = line
		lineCount++
	}
	newJson, err := json.Marshal(newMap)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Generate Folder in root
	workingDir, _ := os.Getwd()
	newFolderPath := workingDir + "\\generatedFiles"
	if _, err := os.Stat(newFolderPath); os.IsNotExist(err) {
		if err := os.Mkdir(newFolderPath, 0755); err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
	}

	// Create File
	fileName := splittedFileName[0]
	fileExtension := ".json"
	newFilePath := newFolderPath + "\\" + fileName + fileExtension
	file, err = os.Create(newFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Write New Values inside File
	val, err := file.Write(newJson)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(val, "successful, check the generatedJsons folder.")
}
