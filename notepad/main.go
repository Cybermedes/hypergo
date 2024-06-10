package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Start with an empty notepad
	notepad := make([]string, 0)
	wordScanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a command and data: ")
		_ = wordScanner.Scan()
		input := strings.Split(wordScanner.Text(), " ")
		command, data := input[0], strings.Join(input[1:], " ")

		switch command {
		case "create":
			create(&notepad, data)
		case "list":
			list(notepad)
		case "clear":
			clear(&notepad)
		case "exit":
			fmt.Println("[INFO] Bye!")
			os.Exit(0)
		default:
			fmt.Println(command)
		}
		fmt.Println()
	}
}

func create(notepad *[]string, note string) {
	// Notepad is set to store up to 5 notes in memory
	if len(*notepad) < 5 {
		*notepad = append(*notepad, note)
		fmt.Println("[OK] The note was successfully created")
	} else {
		fmt.Println("[Error] Notepad is full")
	}
}

func clear(notepad *[]string) {
	*notepad = make([]string, 0)
	fmt.Println("[OK] All notes were successfully deleted")
}

func list(notepad []string) {
	if len(notepad) == 0 {
		return
	} else {
		for i, value := range notepad {
			fmt.Printf("[Info] %d: %s\n", i+1, value)
		}
	}
}
