package values

import "fmt"

/*
Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.

Strings, which can be added together with +.
Integers and floats.
Booleans, with boolean operators as you’d expect.
*/
func PrintValues() {
	fmt.Println(`fmt.Println("go" + "lang") = ` + "go" + "lang")
	fmt.Println(`fmt.Println(1+1) = `, 1+1)
	fmt.Println(`fmt.Println(7.0 / 3.0) = `, 7.0/3.0)
	fmt.Println(`fmt.Println(true && false) = `, true && false)
	fmt.Println(`fmt.Println(true || false) = `, true || false)
	fmt.Println(`fmt.Println(!true) = `, !true)
}
