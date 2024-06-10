package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func create(notepad *[]string, note string, limit int) {

	// Notepad is set to store up to n notes in memory set initially by the user input
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

func update(notepad *[]string, data string) {

	// First, check if notepad is empty
	if isEmpty(notepad) {
		fmt.Println("[Error] There is nothing to update")
		return
	}

	indexAndData := strings.Split(data, " ")

	// Check if there is a note argument
	if len(indexAndData) == 1 {
		fmt.Println("[Error] Missing note argument")
		return
	}

	// Convert the position value from string to int
	index, err := convertStringPosition(indexAndData[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	// The user index is from 1 to n, while the program works with indexes of 0 to n-1
	sliceIndex := index - 1

	// Update the specific note, if it's inside the boundaries
	if _, err := isIndexInRange(sliceIndex, len(*notepad)); err == nil {
		data = strings.Join(indexAndData[1:], " ")
		(*notepad)[sliceIndex] = data
		fmt.Printf("[OK] The note at position %d was successfully updated\n", index)
	} else {
		fmt.Println(err)
	}

}

func delete(notepad *[]string, data string) {

	// First, check if notepad is empty
	if isEmpty(notepad) {
		fmt.Println("[Error] There is nothing to delete")
		return
	}

	// Check if there is a position argument
	if len(data) == 0 {
		fmt.Println("[Error] Missing position argument")
		return
	}

	// Convert the position value from string to int
	index, err := convertStringPosition(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// The user index is from 1 to n, while the program works with indexes of 0 to n-1
	sliceIndex := index - 1

	// Delete the specific note, if it's inside the boundaries
	if _, err := isIndexInRange(sliceIndex, len(*notepad)); err == nil {
		*notepad = append((*notepad)[:sliceIndex], (*notepad)[sliceIndex+1:]...)
		fmt.Printf("[OK] The note at position %d was successfully deleted\n", index)
	} else {
		fmt.Println(err)
	}
}

func isEmpty(notepad *[]string) bool {
	return len(*notepad) == 0
}

func convertStringPosition(index string) (int, error) {
	idx, err := strconv.Atoi(index)
	if err != nil {
		return 0, fmt.Errorf("[Error] Invalid position: %s", index)
	}
	return idx, nil
}

func isIndexInRange(sliceIndex, limit int) (bool, error) {
	if sliceIndex >= 0 && sliceIndex < limit {
		return true, nil
	} else {
		return false, fmt.Errorf("[Error] Position %d is out of the boundaries [1, %d]",
			sliceIndex+1,
			limit+1)
	}
}
