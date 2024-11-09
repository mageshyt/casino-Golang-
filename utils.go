package main  

import (
	"fmt"
	"math/rand"
)
func GetName() string {
	// Quick variable declaration
	// name:= ""

	var name string
	fmt.Println("Welcome To Magesh's Casino...")

	fmt.Println("Enter your name: ")

	_, err := fmt.Scanln(&name)

	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}

	return name
}

func GetBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("Enter Your bet, or o to quite (balance=$%d):", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("You can't bet more than you have")
		} else {
			break
		}
	}

	return bet

}

func GenerateWheel(symbols map[string]uint) []string {
	wheel := []string{}

	for symbol, weight := range symbols {
		for i := uint(0); i < weight; i++ {
			wheel = append(wheel, symbol)
		}

	}
	return wheel
}

func GetRandomNumber(min int, max int) int {
	return rand.Intn(max-min+1) + min
}



func CheckWin(spin [][]string, multiplier map[string]uint) []uint {

	lines := []uint{}

	for _, row := range spin {
		win := true

		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if symbol != checkSymbol {
				win = false
				break
			}

		}

		if win {
			lines = append(lines, multiplier[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}

	return lines

}
