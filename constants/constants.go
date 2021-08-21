package constants

import (
	"fmt"
	"math"
)

const s string = "constant"

func PrintConstants() {
	fmt.Println("\n-------------constants------------")
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
	fmt.Println("------------------------------------")
}
