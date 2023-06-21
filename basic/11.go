// Pointer
package main

import "fmt"

func main() {
	myMsg := "1"
	fmt.Println(myMsg)
	changeMsg(&myMsg, "2")
	fmt.Println(myMsg)
	// You see the difference between the two messages and the original message
	// Only change value of the original message
}

// this is void no return just change value of pointer
func changeMsg(msg *string, newMsg string) {
	*msg = newMsg
}
