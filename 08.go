// While Loop but its For because the go lang does not have while loop.
package main

import "fmt"

func main() {
	x := 0
	for true {
		x++
		if x > 10 {
			break
		}
		fmt.Println(x)
	}
}
