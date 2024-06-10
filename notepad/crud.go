package main

import (
	"fmt"
	"strings"
)

// Create a new note
func create(notepad *[]string, note string, limit int) {

	// Notepad is set to store up to n notes in memory set initially by the user input
	if len(*notepad) < limit {
		*notepad = append(*notepad, note)
		fmt.Println("[OK] The note was successfully created")
	} else {
		fmt.Println("[Error] Notepad is full")
	}
}

// Update a specific note
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

// Delete a specific note
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
