package ranges

import "fmt"

/*
range iterates over elements in a variety of data structures. Let’s see how to use range with some of the data structures we’ve already learned.

Here we use range to sum the numbers in a slice. Arrays work like this too.

range on arrays and slices provides both the index and value for each entry. Above we didn’t need the index, so we ignored it with the blank identifier _. Sometimes we actually want the indexes though.

range on map iterates over key/value pairs.

range can also iterate over just the keys of a map.

range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself.
*/

func PrintRanges() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
