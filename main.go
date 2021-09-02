package main

import (
	"fmt"

	"github.com/well-balanced/gobyexample/morestrings"
	nonblockchanops "github.com/well-balanced/gobyexample/non-blocking-channel-operations"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))

	/* values */
	// values.PrintValues()

	/* variabels */
	// variables.PrintVariables()

	/* constants */
	// constants.PrintConstants()

	/* loop */
	// loop.Loop()

	/* if/else */
	// ifelse.PrintValuesConditionally()

	/* switch */
	// switchexp.Switch()

	/* arrays */
	// arraypkg.PrintArrays()

	/* slices */
	// slices.PrintSlices()

	/* maps */
	// maps.PrintMaps()

	/* ranges */
	// ranges.PrintRanges()

	/* functions */
	// functions.PrintFucntions()

	/* multiple return functions */
	// multiplereturnfunctions.MakeMutipleValueFunction()

	/* variadic functions */
	// variadicfunctions.MakeVariadicFunction()

	/* closures */
	// closures.MakeClosureFunction()

	/* recursion */
	// recursion.RecursiveFunc()

	/* pointers */
	// pointers.UsePointer()

	/* structs */
	// structs.PrintPeople()

	/* methods */
	// methods.CreateMethods()

	/* interfaces */
	// interfaces.Measure()

	/* errors */
	// errortest.ThrowError()

	/* goroutines */
	// goroutines.UseGoroutine()

	/* channels */
	// channels.SendMessage()

	/* channel buffering */
	// chanbuf.ChannelBuffering()

	/* channel synchronization */
	// chansync.SynchronizeChannel()

	/* channel direction */
	// chanway.PingPongMsg()

	/* select */
	// selection.Select()

	/* timeouts */
	// timeouts.Timeout()

	/* non-blocking channel operations */
	nonblockchanops.NonBlockingChannelOperations()
}
