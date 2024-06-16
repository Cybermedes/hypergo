package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Cinema dimensions and matrix of the rows and seats
	var rows, seats int
	fmt.Println("Enter the number of rows:")
	_, _ = fmt.Scan(&rows)
	fmt.Println("Enter the number of seats in each row:")
	_, _ = fmt.Scan(&seats)
	cinema := generateCinema(rows, seats)

	// Run the program in a loop with an 'exit' option
	for {
		var option int
		fmt.Println()
		fmt.Println("1. Show the seats")
		fmt.Println("2. Buy a ticket")
		fmt.Println("0. Exit")
		_, _ = fmt.Scan(&option)

		switch option {
		case 1:
			PrintCinemaLayout(cinema)
		case 2:
			buyTicket(rows, seats, &cinema)
		case 0:
			return
		default:
			fmt.Println("invalid option")
		}
	}

}

func generateCinema(rows, columns int) [][]string {

	cinema := make([][]string, rows+1)
	for i := range cinema {
		cinema[i] = make([]string, columns+1)
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
	fmt.Printf("\nCinema:\n")
	for i := 0; i < len(cinema); i++ {
		for j := 0; j < len(cinema[i]); j++ {
			fmt.Printf("%s ", cinema[i][j])
		}
		fmt.Println()
	}
}

func CalculateMaximumIncome(rows, seats int) int {
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

func buyTicket(rows, seats int, cine *[][]string) {
	var rowNumber, seatNumber int
	fmt.Println()
	fmt.Println("Enter a row number:")
	_, _ = fmt.Scan(&rowNumber)
	fmt.Println("Enter a seat number in that row:")
	_, _ = fmt.Scan(&seatNumber)

	(*cine)[rowNumber][seatNumber] = "B"

	fmt.Printf("Ticket price: $%d\n", checkPrice(rows, seats, rowNumber))

}

func checkPrice(rows, seats, rowNumber int) int {
	if rows*seats < 60 {
		return 10
	} else {
		if rowNumber <= rows/2 {
			return 10
		} else {
			return 8
		}
	}
}
