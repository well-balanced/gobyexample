package values

import "fmt"

func PrintValues() {
	fmt.Println("\n-------------values------------")
	fmt.Println(`fmt.Println("go" + "lang") = ` + "go" + "lang")
	fmt.Println(`fmt.Println(1+1) = `, 1+1)
	fmt.Println(`fmt.Println(7.0 / 3.0) = `, 7.0/3.0)
	fmt.Println(`fmt.Println(true && false) = `, true && false)
	fmt.Println(`fmt.Println(true || false) = `, true || false)
	fmt.Println(`fmt.Println(!true) = `, !true)
	fmt.Println("-------------------------------")
}
