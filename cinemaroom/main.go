package main

import (
	"fmt"
	"strconv"
)

func main() {
	var rows, seats int
	fmt.Println("Enter the number of rows:")
	_, _ = fmt.Scan(&rows)
	fmt.Println("Enter the number of seats in each row:")
	_, _ = fmt.Scan(&seats)

	income := calculateMaximumIncome(rows, seats)
	fmt.Println("Total income:")
	fmt.Printf("$%d\n", income)
}

func GenerateCinema(rows, columns int) [][]string {

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

func PrintCinemaLayout(cinema [][]string) {
	for i := 0; i < len(cinema); i++ {
		for j := 0; j < len(cinema[i]); j++ {
			fmt.Printf("%s ", cinema[i][j])
		}
		fmt.Println()
	}
}

func calculateMaximumIncome(rows, seats int) int {
	if rows*seats < 60 {
		return rows * seats * 10
	} else {
		if rows%2 == 0 {
			return 8*(rows/2)*seats + 10*(rows/2)*seats
		} else {
			return 8*(rows/2+1)*seats + 10*(rows/2)*seats
		}
	}
}
