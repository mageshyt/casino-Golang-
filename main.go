package main

import (
	"fmt"
)


func main() {

	symbols := map[string]uint{
		"🍒": 4,
		"🍋": 7,
		"🍊": 12,
		"🍇": 20,
		"🍉": 25,
		"🍎": 30,
		"🍓": 32,
	}

	// multiplier to multiply the bet
	multiplier := map[string]uint{
		"🍒": 20,
		"🍋": 10,
		"🍊": 8,
		"🍇": 5,
		"🍉": 6,
		"🍎": 3,
		"🍓": 2,
	}

	wheel := GenerateWheel(symbols)

	GetName()
	balance := uint(200)

	for balance > 0 {
		bet := GetBet(balance)

		// BASE CASE
		if bet == 0 {
			break
		}
		// BALANCE = BALANCE - BET
		balance -= bet
		spin := SpinWheel(wheel, 3, 3)

		PrintSpin(spin)

		winningLines := CheckWin(spin, multiplier)

		totalWin := uint(0)

		for i, multi := range winningLines {
			win := multi * bet
			totalWin += win
			if win > 0 {

				fmt.Printf("Won $%d, (%dx) on Line #%d \n", win, multi, i+1)
			}

		}

		if totalWin > 0 {
			// BALANCE = BALANCE + TOTALWIN
			balance += totalWin
		}
	}

	fmt.Printf("You left with $ %d \n", balance)

}
