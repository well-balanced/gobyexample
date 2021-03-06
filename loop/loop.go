package loop

import "fmt"

/*
for is Go’s only looping construct. Here are some basic types of for loops.

The most basic type, with a single condition.

A classic initial/condition/after for loop.

for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.

You can also continue to the next iteration of the loop.
*/

func Loop() {
	defer fmt.Println()

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
