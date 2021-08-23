package main

import (
	"fmt"

	"github.com/well-balanced/gobyexample/arraypkg"
	"github.com/well-balanced/gobyexample/constants"
	"github.com/well-balanced/gobyexample/functions"
	"github.com/well-balanced/gobyexample/ifelse"
	"github.com/well-balanced/gobyexample/loop"
	"github.com/well-balanced/gobyexample/maps"
	"github.com/well-balanced/gobyexample/morestrings"
	"github.com/well-balanced/gobyexample/ranges"
	"github.com/well-balanced/gobyexample/slices"
	"github.com/well-balanced/gobyexample/switchexp"
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

	/* loop */
	loop.Loop()

	/* if/else */
	ifelse.PrintValuesConditionally()

	/* switch */
	switchexp.Switch()

	/* array */
	arraypkg.PrintArrays()

	/* slices */
	slices.PrintSlices()

	/* maps */
	maps.PrintMaps()

	/* ranges */
	ranges.PrintRanges()

	/* functions */
	functions.PrintFucntions()
}
