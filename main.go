package main

import (
	"fmt"

	"github.com/well-balanced/gobyexample/arraypkg"
	"github.com/well-balanced/gobyexample/chanbuf"
	"github.com/well-balanced/gobyexample/channels"
	"github.com/well-balanced/gobyexample/closures"
	"github.com/well-balanced/gobyexample/constants"
	"github.com/well-balanced/gobyexample/errortest"
	"github.com/well-balanced/gobyexample/functions"
	"github.com/well-balanced/gobyexample/goroutines"
	"github.com/well-balanced/gobyexample/ifelse"
	"github.com/well-balanced/gobyexample/interfaces"
	"github.com/well-balanced/gobyexample/loop"
	"github.com/well-balanced/gobyexample/maps"
	"github.com/well-balanced/gobyexample/methods"
	"github.com/well-balanced/gobyexample/morestrings"
	"github.com/well-balanced/gobyexample/multiplereturnfunctions"
	"github.com/well-balanced/gobyexample/pointers"
	"github.com/well-balanced/gobyexample/ranges"
	"github.com/well-balanced/gobyexample/recursion"
	"github.com/well-balanced/gobyexample/slices"
	"github.com/well-balanced/gobyexample/structs"
	"github.com/well-balanced/gobyexample/switchexp"
	"github.com/well-balanced/gobyexample/values"
	"github.com/well-balanced/gobyexample/variables"
	"github.com/well-balanced/gobyexample/variadicfunctions"
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

	/* arrays */
	arraypkg.PrintArrays()

	/* slices */
	slices.PrintSlices()

	/* maps */
	maps.PrintMaps()

	/* ranges */
	ranges.PrintRanges()

	/* functions */
	functions.PrintFucntions()

	/* multiple return functions */
	multiplereturnfunctions.MakeMutipleValueFunction()

	/* variadic functions */
	variadicfunctions.MakeVariadicFunction()

	/* closures */
	closures.MakeClosureFunction()

	/* recursion */
	recursion.RecursiveFunc()

	/* pointers */
	pointers.UsePointer()

	/* structs */
	structs.PrintPeople()

	/* methods */
	methods.CreateMethods()

	/* interfaces */
	interfaces.Measure()

	/* errors */
	errortest.ThrowError()

	/* goroutines */
	goroutines.UseGoroutine()

	/* channels */
	channels.SendMessage()

	/* channel buffering */
	chanbuf.ChannelBuffering()
}
