package main

import (
	"fmt"
	"strconv"
)

func list(notepad []string) {
	if len(notepad) == 0 {
		fmt.Println("[Info] Notepad is empty")
	} else {
		for i, value := range notepad {
			fmt.Printf("[Info] %d: %s\n", i+1, value)
		}
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
