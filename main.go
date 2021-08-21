package main

import (
	"fmt"

	"github.com/well-balanced/gobyexample/constants"
	"github.com/well-balanced/gobyexample/morestrings"
	"github.com/well-balanced/gobyexample/values"
	"github.com/well-balanced/gobyexample/variables"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))

	/* values */
	values.PrintValues()

	/* variabels */
	variables.PrintVariables()

	/* constants */
	constants.PrintConstants()

}
