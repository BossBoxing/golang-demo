package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Hey")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("I'm alive now!")
}
