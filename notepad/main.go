package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Set a maximum numbers of notes that the notepad can store
	fmt.Print("Enter the maximum number of notes: ")
	var numberOfNotes int
	_, _ = fmt.Scan(&numberOfNotes)
	fmt.Println()

	// Start with an empty notepad
	notepad := make([]string, 0)
	wordScanner := bufio.NewScanner(os.Stdin)

	// REPL(Read, Evaluate, Print, Loop) mode
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
		case "update":
			if len(data) == 0 {
				fmt.Println("[Error] Missing position argument")
			} else {
				update(&notepad, data)
			}
		case "delete":
			delete(&notepad, data)
		case "exit":
			fmt.Println("[INFO] Bye!")
			os.Exit(0)
		default:
			fmt.Println("[Error] Unknown command")
		}
		fmt.Println()
	}
}

func clear(notepad *[]string) {
	*notepad = make([]string, 0)
	fmt.Println("[OK] All notes were successfully deleted")
}
