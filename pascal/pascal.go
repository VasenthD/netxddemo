package main

import "fmt"

func main() {
	fmt.Println("ENTER THE NUMBER OF ROWS :")
	var input int
	fmt.Scanf("%d", &input)
	var pascal = make([][]int, input)
	for i := 0; i < input; i++ {
		pascal[i] = make([]int, input)
	}

	for i := 0; i < input; i++ {
		pascal[i][0] = 1
		pascal[i][i] = 1
		var result int = 1
		var n int = i
		for j := 1; j < i; j++ {
			result = result * n
			result = result / j
			pascal[i][j] = result
			n--
		}
	}
	for i := 0; i < input; i++ {
		x := input - i
		for j := 0; j <= input-i; j++ {
			if j == x {
				for k := 0; k <= i; k++ {
					fmt.Print(pascal[i][k], " ")
				}
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println(" ")
	}
}
