package chanway

import "fmt"

/*
When using channels as function parameters,
you can specify if a channel is meant to only send or receive values.
This specificity increases the type-safety of the program.
*/

func ping(pings chan<- string, msg string) {
	/*
		This ping function only accepts a channel for sending values.
		It would be a compile-time error to try to receive on this channel.
	*/
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	/*
		The pong function accepts one channel for receives (pings) and a second for sends (pongs).
	*/
	msg := <-pings
	pongs <- msg
}

func PingPongMsg() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "hello, strangers.")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
