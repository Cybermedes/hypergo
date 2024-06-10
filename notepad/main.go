package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	wordScanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a command and data: ")
		_ = wordScanner.Scan()
		input := strings.Split(wordScanner.Text(), " ")
		command, data := input[0], strings.Join(input[1:], " ")

		switch command {
		case "create":
			fmt.Println(command, data)
		case "list":
			fmt.Println("list")
		case "exit":
			fmt.Println("[INFO] Bye!")
			os.Exit(0)
		default:
			fmt.Println(command)
		}
		fmt.Println()
	}
}
