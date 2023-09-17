package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text2jsonConverter/helpers"
)

func main() {
	for {
		fmt.Println("")
		fmt.Println("Welcome to the file converter🍾")
		fmt.Println("Ensure you're in the right directory then input your command:")
		fmt.Println("eg: text2json file_name.txt")

		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Fields(input)

		command, fileName := args[0], args[1]

		if len(args) != 2 {
			fmt.Println("Kindly input the right commands!")
			return
		}

		switch command {
		case "text2json":
			path, splittedFileNameArr := helpers.ValidateAndLocateFile(fileName)
			helpers.ConvertText2Json(path, splittedFileNameArr)
		default:
			fmt.Println("Invalid command. Please try again.")
			continue
		}

		break
	}
}
