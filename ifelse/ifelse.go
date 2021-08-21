package ifelse

import "fmt"

func PrintValuesConditionally() {
	fmt.Println("\n-------------if/else------------")
	testNumber := 7
	if testNumber%2 == 0 {
		fmt.Println(testNumber, "is even")
	} else {
		fmt.Println(testNumber, "is odd")
	}

	testNumber = 8
	if testNumber%4 == 0 {
		fmt.Println(testNumber, "is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	fmt.Println("-------------------------------")
}
