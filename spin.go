package main

import ("fmt")

func SpinWheel(wheel []string, rows int, cols int) [][]string {
	// Randomly select 3 symbols from the wheel
	result := [][]string{}

	for i := 0; i < rows; i++ {
		// add the empty array to the result
		result = append(result, []string{})
	}

	for col := 0; col < cols; col++ {
		selected := map[int]bool{}
		for row := 0; row < rows; row++ {
			// get a random number until we get a unique number
			for true {
				randomNumber := GetRandomNumber(0, len(wheel)-1)

				// check if the random number is already selected
				_, exists := selected[randomNumber]

				if !exists {
					selected[randomNumber] = true
					result[row] = append(result[row], wheel[randomNumber])
					break
				}
			}

		}

	}
	return result

}

func PrintSpin(spin [][]string) {
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Printf(symbol)
			if j != len(row)-1 {
				fmt.Printf(" | ")
			}

		}
		fmt.Println()
	}
}
