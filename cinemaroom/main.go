package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Cinema dimensions and matrix of the rows and seats
	var rows, seats, balance, ticketsSold int
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
		fmt.Println("3. Statistics")
		fmt.Println("0. Exit")
		_, _ = fmt.Scan(&option)

		switch option {
		case 1:
			printCinemaLayout(cinema)
		case 2:
			buyTicket(rows, seats, &cinema, &balance, &ticketsSold)
		case 3:
			showStatistics(rows, seats, &ticketsSold, &balance)
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

func printCinemaLayout(cinema [][]string) {
	fmt.Printf("\nCinema:\n")
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

func buyTicket(rows, seats int, cine *[][]string, balance, ticket *int) {

	for {
		var rowNumber, seatNumber int
		fmt.Println()
		fmt.Println("Enter a row number:")
		_, _ = fmt.Scan(&rowNumber)
		fmt.Println("Enter a seat number in that row:")
		_, _ = fmt.Scan(&seatNumber)

		if rowNumber > rows || seatNumber > seats {
			fmt.Printf("\nWrong input!\n")
			continue
		}

		if (*cine)[rowNumber][seatNumber] == "B" {
			fmt.Printf("\nThat ticket has already been purchased!\n")
		} else {
			(*cine)[rowNumber][seatNumber] = "B"
			ticketPrice := checkSeatPrice(rows, seats, rowNumber)
			*balance += ticketPrice
			*ticket += 1
			fmt.Printf("\nTicket price: $%d\n", ticketPrice)
			return
		}
	}
}

func checkSeatPrice(rows, seats, rowNumber int) int {
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

func showStatistics(rows, seats int, ticketsSold, balance *int) {

	occupancy := (float64(*ticketsSold) / float64(rows*seats)) * 100
	fmt.Println()
	fmt.Printf("Number of purchased tickets: %d\n", *ticketsSold)
	fmt.Printf("Percentage: %.2f%%\n", occupancy)
	fmt.Printf("Current income: $%d\n", *balance)
	fmt.Printf("Total income: $%d\n", calculateMaximumIncome(rows, seats))
}
