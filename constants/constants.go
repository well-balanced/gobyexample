package constants

import (
	"fmt"
	"math"
)

/*
Go supports constants of character, string, boolean, and numeric values.

const declares a constant value.

A const statement can appear anywhere a var statement can.

Constant expressions perform arithmetic with arbitrary precision.

A numeric constant has no type until it’s given one, such as by an explicit conversion.

A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.
*/

const s string = "constant"

func PrintConstants() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
