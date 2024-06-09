package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Cinema:")
	room := generateCinema(8, 9)
	printCinemaLayout(room)
}

func generateCinema(rows, columns int) [][]string {

	cinema := make([][]string, rows)
	for i := range cinema {
		cinema[i] = make([]string, columns)
		for j := range cinema[i] {
			if i == 0 {
				cinema[i][j] = strconv.Itoa(j)
			} else if j == 0 {
				cinema[i][j] = strconv.Itoa(i)
			} else {
				cinema[i][j] = "S"
			}
		}
	}
	cinema[0][0] = " "
	return cinema
}

func printCinemaLayout(cinema [][]string) {
	for i := 0; i < len(cinema); i++ {
		for j := 0; j < len(cinema[i]); j++ {
			fmt.Printf("%s ", cinema[i][j])
		}
		fmt.Println()
	}
}
