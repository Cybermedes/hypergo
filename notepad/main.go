package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter the maximum number of notes: ")
	var numberOfNotes int
	_, _ = fmt.Scan(&numberOfNotes)
	fmt.Println()

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
			if len(data) == 0 {
				fmt.Println("[Error] Missing note argument")
			} else {
				create(&notepad, data, numberOfNotes)
			}
		case "list":
			list(notepad)
		case "clear":
			clear(&notepad)
		case "exit":
			fmt.Println("[INFO] Bye!")
			os.Exit(0)
		default:
			fmt.Println("[Error] Unknown command")
		}
		fmt.Println()
	}
}

func create(notepad *[]string, note string, limit int) {

	// Notepad is set to store up to 5 notes in memory
	if len(*notepad) < limit {
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
		fmt.Println("[Info] Notepad is empty")
	} else {
		for i, value := range notepad {
			fmt.Printf("[Info] %d: %s\n", i+1, value)
		}
	}
}
